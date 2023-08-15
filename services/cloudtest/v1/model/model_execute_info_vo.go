package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteInfoVo 执行信息
type ExecuteInfoVo struct {

	// 执行开始时间
	Time *string `json:"time,omitempty"`

	// 执行开始时间时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 执行时长
	Duration *string `json:"duration,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 执行次数
	ExecuteTimes *int32 `json:"execute_times,omitempty"`
}

func (o ExecuteInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteInfoVo struct{}"
	}

	return strings.Join([]string{"ExecuteInfoVo", string(data)}, " ")
}
