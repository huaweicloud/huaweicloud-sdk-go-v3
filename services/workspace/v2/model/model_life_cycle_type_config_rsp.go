package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LifeCycleTypeConfigRsp 生命周期类型配置响应。
type LifeCycleTypeConfigRsp struct {

	// 最小等待时长，单位分钟。
	MinWaitTime *int32 `json:"min_wait_time,omitempty"`

	// 默认等待时长，单位分钟。
	DefaultWaitTime *int32 `json:"default_wait_time,omitempty"`

	// 最小执行周期，单位分钟。
	MinExecTime *int32 `json:"min_exec_time,omitempty"`

	// 默认执行周期，单位分钟。
	DefaultExecTime *int32 `json:"default_exec_time,omitempty"`

	// 是否支持执行周期。
	SupportExecTime *bool `json:"support_exec_time,omitempty"`

	// 可执行的动作列表。
	Actions *[]ActionConfig `json:"actions,omitempty"`
}

func (o LifeCycleTypeConfigRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LifeCycleTypeConfigRsp struct{}"
	}

	return strings.Join([]string{"LifeCycleTypeConfigRsp", string(data)}, " ")
}
