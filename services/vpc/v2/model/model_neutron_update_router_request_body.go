package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateRouterRequestBody
type NeutronUpdateRouterRequestBody struct {
	Router *NeutronUpdateRouterOption `json:"router"`
}

func (o NeutronUpdateRouterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateRouterRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateRouterRequestBody", string(data)}, " ")
}
