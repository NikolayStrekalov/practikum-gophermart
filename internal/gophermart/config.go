package gophermart

type Config struct {
	SelfAddress    string
	Database       string
	AccrualAddress string
}

func GetConfig() Config {
	return Config{
		SelfAddress: "localhost:8080",
	}
}
