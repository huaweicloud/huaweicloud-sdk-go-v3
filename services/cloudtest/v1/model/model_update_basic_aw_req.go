package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBasicAwReq struct {

	// 描述
	Description *string `json:"description,omitempty"`

	// 默认检查点List
	DftCheckPointList *[]CheckPoint `json:"dft_check_point_list,omitempty"`

	// AW参数类list
	DftCustomHeader *[]AwParam `json:"dft_custom_header,omitempty"`

	// 重试间隔时间 (ms) 为空表示不等待(目前内部使用)
	DftRetryInterval *string `json:"dft_retry_interval,omitempty"`

	// 重试次数(目前内部使用)
	DftRetryTimes *string `json:"dft_retry_times,omitempty"`

	// 定义的变量信息
	DftVariableList *[]AwVariable `json:"dft_variable_list,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 参数类型和参数默认值对应List
	ParamTypeAndDftValue *[]AwParam `json:"param_type_and_dft_value,omitempty"`
}

func (o UpdateBasicAwReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBasicAwReq struct{}"
	}

	return strings.Join([]string{"UpdateBasicAwReq", string(data)}, " ")
}
