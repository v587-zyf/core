package db

type Config struct {
	Uri string `json:"uri"`
	DB  string `json:"db"`
	Use bool   `json:"use"`
}
