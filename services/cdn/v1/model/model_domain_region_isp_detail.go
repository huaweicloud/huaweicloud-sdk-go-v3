package model

import (
	"encoding/json"

	"strings"
)

type DomainRegionIspDetail struct {
	// 区域

	Region *string `json:"region,omitempty"`
	// 运营商

	Isp *string `json:"isp,omitempty"`
	// 流量

	Flux *[]int32 `json:"flux,omitempty"`
}

func (o DomainRegionIspDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DomainRegionIspDetail struct{}"
	}

	return strings.Join([]string{"DomainRegionIspDetail", string(data)}, " ")
}
