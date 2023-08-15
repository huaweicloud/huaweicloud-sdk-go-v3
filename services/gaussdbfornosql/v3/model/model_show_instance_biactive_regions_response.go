package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceBiactiveRegionsResponse Response Object
type ShowInstanceBiactiveRegionsResponse struct {
	RegionCodes    *[]string `json:"region_codes,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowInstanceBiactiveRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceBiactiveRegionsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceBiactiveRegionsResponse", string(data)}, " ")
}
