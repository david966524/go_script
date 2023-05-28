package cf

import (
	"context"
	"fmt"
	"log"
)

var (
	key   string = ""
	email string = ""
)

type CoudFlare struct {
	Domain   string
	OriginIp string
	CnameUrl string
}

func (cf *CoudFlare) CreateDomain() {
	api, err := cfapi.New(key, email)
	if err != nil {
		log.Fatal(err)
	}
	zoneParams := cfapi.ZoneCreateParams{
		Name:      cf.Domain,
		JumpStart: false,
	}
	zone, err := api.CreateZone(context.Background(), zoneParams.Name, zoneParams.JumpStart, cfapi.Account{}, "")

	if err != nil {
		fmt.Println(err.Error())
	}

}

func (cf *CoudFlare) AddRecords() {

}
