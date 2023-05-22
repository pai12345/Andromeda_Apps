package regions

import (
	"errors"
	"fmt"
)

func FetchAllregions() *SupportedRegions {
	supportedregions := SupportedRegions{
		Regionsupport: map[string]string{
			"Europe":       "EU",
			"NorthAmerica": "US",
			"EMEA":         "ALL",
		},
	}
	fmt.Println(&supportedregions)
	return &supportedregions
}

func (regionsupport SupportedRegions) CheckRegionsupport(region string) error {
	_, check := regionsupport.Regionsupport[region]
	if !check {
		return errors.New("unsupported region")
	}
	return nil
}

func (regionsupport SupportedRegions) FetchRegion(region string) map[string]string {
	result := map[string]string{}
	for key, value := range regionsupport.Regionsupport {
		if region == key {
			result[key] = value
			break
		}
	}
	fmt.Println(result)
	return result
}
