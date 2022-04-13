package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAccessoryLimitsResponse struct {
	AccessoryLimit *AccessoryLimitVo `json:"accessory_limit,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowAccessoryLimitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessoryLimitsResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessoryLimitsResponse", string(data)}, " ")
}
