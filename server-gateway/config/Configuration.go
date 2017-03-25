package config

import (
	"time"

	arg "github.com/alexflint/go-arg"
)

var RawEnv struct {
	Mode string `arg:"env"`
	Port string `arg:"env"`
	Brokers string `arg:"env"`
}

type Environment struct {
	Mode string
	Port string
	Brokers string
}

func GetEnvironment() Environment {
	// Setup default values
	RawEnv.Mode = "LOCAL" // [LOCAL, DEV, PROD]
	RawEnv.Port = "10001"
	RawEnv.Brokers = "localhost:9092"
	arg.MustParse(&RawEnv)

	return Environment{
		Mode: RawEnv.Mode,
		Port: ":" + RawEnv.Port,
		Brokers: RawEnv.Brokers,
	}
}

var (
	Version   = "undefined"
	BuildTime = "undefined"
	GitHash   = "undefined"
	Started   = time.Now()
)

type Flag struct {
	Version   string
	BuildTime string
	GitHash   string
	Started   string
}

func GetFlag() Flag {
	return Flag{
		Version:   Version,
		BuildTime: BuildTime,
		GitHash:   GitHash,
		Started:   Started.UTC().Format(time.RFC3339),
	}
}
