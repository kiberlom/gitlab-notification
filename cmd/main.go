package main

import (
	"log"

	"github.com/kiberlom/gitlab-notifications/internal/config"
	v1 "github.com/kiberlom/gitlab-notifications/internal/handlers/http/v1"
	"github.com/kiberlom/gitlab-notifications/internal/servers"
	"github.com/kiberlom/gitlab-notifications/internal/shutdown"
)

func main() {
	cfgENV := config.NewENV()
	_ = config.NewConfig()

	g, _ := shutdown.NewShutdown()

	muxV1 := v1.NewHandler(&v1.HandlerSetup{})

	g.Go(func() error {
		if err := servers.NewHTTPServer(cfgENV, muxV1); err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("service stop ...")
}
