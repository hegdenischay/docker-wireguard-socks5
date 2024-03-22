package main

import (
	"log"
	// "net"
	"os"
	// "os/exec"

	// "github.com/armon/go-socks5"
	"github.com/wzshiming/socks5"
	"github.com/caarlos0/env"
)

type params struct {
	User     string `env:"SOCKS5_USER" envDefault:""`
	Password string `env:"SOCKS5_PASS" envDefault:""`
	Port     string `env:"SOCKS5_PORT" envDefault:"1080"`
	Up       string `env:"SOCKS5_UP"   envDefault:""`
}

func main() {
	// Working with app params
	cfg := params{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	// initialize logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	//Initialize socks5 config
	socks5conf := &socks5.Server{
		Logger: logger,
	}

	if cfg.User+cfg.Password != "" {
		// creds := socks5.StaticCredentials{
		// 	cfg.User: cfg.Password,
		// }
		// cator := socks5.UserPassAuthenticator{Credentials: creds}
		socks5conf.Authentication = socks5.UserAuth(cfg.User, cfg.Password)
	}

	err = socks5conf.ListenAndServe("tcp", ":"+cfg.Port)
	if err != nil {

		logger.Println(err)
	}
}
