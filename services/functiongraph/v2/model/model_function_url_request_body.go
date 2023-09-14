package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FunctionUrlRequestBody struct {

	// 函数URL鉴权方式
	AuthType string `json:"auth_type"`

	Cors *CorsConfig `json:"cors"`
}

func (o FunctionUrlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionUrlRequestBody struct{}"
	}

	return strings.Join([]string{"FunctionUrlRequestBody", string(data)}, " ")
}
