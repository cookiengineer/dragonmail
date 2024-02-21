package structs

type Config struct {
	Domain  string `json:"domain"`
	Formats struct {
		Login string `json:"login"`
	} `json:"formats,omitempty"`
	Protocols []Protocol `json:"protocols"`
}
