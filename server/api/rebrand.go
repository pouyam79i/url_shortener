package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/config"
	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/util"
)

// build and send a request to rebrandly, then return result
func CallRebrandlyAPI(url string) (string, error) {

	domainProp := config.DomainProp{
		FullName: "rebrand.ly",
	}
	body := config.RebrandlyRequestAPI{
		Destination: url,
		Domain:      domainProp,
	}

	byteBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// loading configs:
	server_conf, err := util.GetConfigs()

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, server_conf.RebrandlyURL, bytes.NewBuffer(byteBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("apikey", server_conf.API_KEY)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	var data map[string]interface{}
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		return "", err
	}

	if data["shortUrl"] == nil {
		return "", errors.New("shortened url does not exists in response from rebrandly api")
	}

	shortUrl := fmt.Sprint(data["shortUrl"])
	return shortUrl, nil

}
