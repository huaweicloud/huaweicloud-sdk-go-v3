package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SystemParametersResult struct {

	// 编号
	Id *int32 `json:"id,omitempty"`

	// 参数名称
	ParameterName *string `json:"parameter_name,omitempty"`

	// 参数描述
	Description *string `json:"description,omitempty"`
}

func (o SystemParametersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemParametersResult struct{}"
	}

	return strings.Join([]string{"SystemParametersResult", string(data)}, " ")
}
