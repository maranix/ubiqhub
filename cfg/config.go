package cfg

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"post"`
}
