package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableResourceResp struct {
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Flavors *[]FlavorResource `json:"flavors,omitempty"`
}

func (o AvailableResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableResourceResp struct{}"
	}

	return strings.Join([]string{"AvailableResourceResp", string(data)}, " ")
}
