package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreatePortRequestBody
type NeutronCreatePortRequestBody struct {
	Port *NeutronCreatePortOption `json:"port"`
}

func (o NeutronCreatePortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreatePortRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreatePortRequestBody", string(data)}, " ")
}
