package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateNetworkRequestBody
type NeutronUpdateNetworkRequestBody struct {
	Network *NeutronUpdateNetworkOption `json:"network"`
}

func (o NeutronUpdateNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateNetworkRequestBody", string(data)}, " ")
}
