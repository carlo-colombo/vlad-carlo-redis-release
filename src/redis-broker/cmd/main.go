package main

import (
	"net/http"
	"redis-broker"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

func main() {

	serviceBroker := broker.New()
	logger := lager.NewLogger("vlad-carlo-redis-broker")
	credentials := brokerapi.BrokerCredentials{
		Username: "username",
		Password: "password",
	}

	brokerAPI := brokerapi.New(serviceBroker, logger, credentials)

	http.Handle("/", brokerAPI)
	http.ListenAndServe(":3000", nil)
}
