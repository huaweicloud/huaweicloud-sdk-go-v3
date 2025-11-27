package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteDocumentRequestBodyTargets struct {

	// 实例化执行目标的维度，枚举：InstanceValues，BatchValues
	Key *string `json:"key,omitempty"`

	// 根据target_parameter_name指定全局参数指定要执行的目标实例
	Values *[]string `json:"values,omitempty"`
}

func (o ExecuteDocumentRequestBodyTargets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequestBodyTargets struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequestBodyTargets", string(data)}, " ")
}
