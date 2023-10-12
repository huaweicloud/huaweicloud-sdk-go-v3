package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvRequestBodyParams struct {

	// 环境变量名称
	Name *string `json:"name,omitempty"`

	// 环境变量类型
	Type *string `json:"type,omitempty"`

	// 环境变量值
	Value *string `json:"value,omitempty"`

	// 描述信息
	Desc *string `json:"desc,omitempty"`
}

func (o EnvRequestBodyParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvRequestBodyParams struct{}"
	}

	return strings.Join([]string{"EnvRequestBodyParams", string(data)}, " ")
}
