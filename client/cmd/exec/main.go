package main

import (
	"cv-client/pkg/api"
	"cv-client/pkg/log"
)

func main() {

	log.Motd()

	sr := api.SecurityReport{}
	sr.MakeSecurityReport()

}
