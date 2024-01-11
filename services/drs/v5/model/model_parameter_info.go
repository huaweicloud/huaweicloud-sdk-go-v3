package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParameterInfo 任务参数配置信息
type ParameterInfo struct {

	// 参数名称，如：“applier_thread_num”，“read_task_num”等
	ParameterName string `json:"parameter_name"`

	// 参数名称对应的参数值，如：“20”，“false”
	ParameterValue string `json:"parameter_value"`
}

func (o ParameterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterInfo struct{}"
	}

	return strings.Join([]string{"ParameterInfo", string(data)}, " ")
}
