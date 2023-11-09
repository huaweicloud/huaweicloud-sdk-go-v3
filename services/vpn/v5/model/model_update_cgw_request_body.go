package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCgwRequestBody struct {
	CustomerGateway *UpdateCgwRequestBodyContent `json:"customer_gateway"`
}

func (o UpdateCgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCgwRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCgwRequestBody", string(data)}, " ")
}
