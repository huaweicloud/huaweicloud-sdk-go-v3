package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdatePortRequestBody
type NeutronUpdatePortRequestBody struct {
	Port *NeutronUpdatePortOption `json:"port"`
}

func (o NeutronUpdatePortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdatePortRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdatePortRequestBody", string(data)}, " ")
}
