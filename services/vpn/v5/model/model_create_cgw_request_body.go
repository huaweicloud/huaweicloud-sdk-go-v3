package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCgwRequestBody struct {
	CustomerGateway *CreateCgwRequestBodyContent `json:"customer_gateway"`
}

func (o CreateCgwRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCgwRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCgwRequestBody", string(data)}, " ")
}
