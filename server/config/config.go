package config

// TODO: remove this hard code add it to env variable!
const (
	Config_Path = "./config.yml"
)

// Server main configs to connect to rebrandly
type Server struct {
	RebrandlyURL string `yaml:"rebrandlyUrl"`
	API_KEY      string `yaml:"apikey"`
	REDIS_ADDR   string `yaml:"redis"`
	REDIS_TIME   string `yaml:"redisTime"`
}

// req for my server
type RequestAPI struct {
	URL string `json:"url"`
}

// res for my server
type ResponseAPI struct {
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
	IsCached bool   `json:"isCached"`
	Hostname string `json:"hostname"`
}

// req for rebrandly
type DomainProp struct {
	FullName string `json:"fullName"`
}
type RebrandlyRequestAPI struct {
	Destination string     `json:"destination"`
	Domain      DomainProp `json:"domain"`
}
