package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateParamInfo struct {

	// 节点类型。取值范围： - be - fe
	NodeType string `json:"node_type"`

	// 参数名和参数值映射关系。用户可以基于默认参数模板的参数，自定义的参数值。不传入该参数，则保持原参数信息。
	ParameterValues map[string]string `json:"parameter_values"`
}

func (o UpdateParamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParamInfo struct{}"
	}

	return strings.Join([]string{"UpdateParamInfo", string(data)}, " ")
}
