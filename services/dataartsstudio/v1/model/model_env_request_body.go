package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvRequestBody struct {

	// 环境变量实体信息
	Params *[]EnvRequestBodyParams `json:"params,omitempty"`
}

func (o EnvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvRequestBody struct{}"
	}

	return strings.Join([]string{"EnvRequestBody", string(data)}, " ")
}
