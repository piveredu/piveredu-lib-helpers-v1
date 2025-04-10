package main

import (
	"github.com/piveredu/piveredu-lib-helpers-v1/pivereduhelpers"
	"log"
	"os"
)

func main() {
	contextHelper := pivereduhelpers.NewContextHelper()
	accessToken, err := contextHelper.GetAccessToken(nil)
	if err != nil {
		os.Exit(1)
	}

	log.Println("AccessToken:", accessToken)

	tenant, err := contextHelper.GetTenant(nil)
	if err != nil {
		os.Exit(1)
	}

	log.Println("Tenant:", tenant)
}
