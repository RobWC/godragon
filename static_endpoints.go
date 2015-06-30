package main

import "fmt"

const (
	StaticScheme      = "https"
	StaticDataVersion = "v1.2"
	StaticAPIPath     = "api/lol/static-data"
	LocalCode         = "en_US"
	RegionNA          = "NA"
	RegionPBE         = "PBE"
)

func staticDataAPI(region string, kind string) string {
	return fmt.Sprintf("/%s/%s/%s/%s", StaticAPIPath, region, StaticDataVersion, kind)
}

func staticDataByIDAPI(region string, kind string, id string) string {
	return fmt.Sprintf("/%s/%s/%s/%s/%s", StaticAPIPath, region, StaticDataVersion, kind, id)
}

func Champions(tversion string) string {
	return fmt.Sprintf("/")
}
