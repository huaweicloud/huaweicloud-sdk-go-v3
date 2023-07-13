package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateRouterRequestBody
type NeutronCreateRouterRequestBody struct {
	Router *NeutronCreateRouterOption `json:"router"`
}

func (o NeutronCreateRouterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateRouterRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateRouterRequestBody", string(data)}, " ")
}
