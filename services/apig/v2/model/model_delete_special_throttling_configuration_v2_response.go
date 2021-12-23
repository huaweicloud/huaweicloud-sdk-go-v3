package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSpecialThrottlingConfigurationV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSpecialThrottlingConfigurationV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSpecialThrottlingConfigurationV2Response struct{}"
	}

	return strings.Join([]string{"DeleteSpecialThrottlingConfigurationV2Response", string(data)}, " ")
}
