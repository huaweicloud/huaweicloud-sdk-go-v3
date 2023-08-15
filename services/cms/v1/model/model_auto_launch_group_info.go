package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoLaunchGroupInfo 智能购买组列表
type AutoLaunchGroupInfo struct {

	// 智能购买组ID
	Id string `json:"id"`

	// 智能购买组的名称
	Name string `json:"name"`

	// 请求类型
	Type string `json:"type"`

	// 智能购买组的运行状态，枚举值 SUBMITTED：已提交 ACTIVE：运行中 DELETING：删除中 DELETED：已删除
	Status string `json:"status"`

	// 智能购买组的任务状态，枚举值 HANDLING：购买中 FULFILLED：智能购买组已满配 ERROR：智能购买组异常
	TaskState string `json:"task_state"`

	// 开始时间
	ValidSince *sdktime.SdkTime `json:"valid_since"`

	// 结束时间
	ValidUntil *sdktime.SdkTime `json:"valid_until"`
}

func (o AutoLaunchGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLaunchGroupInfo struct{}"
	}

	return strings.Join([]string{"AutoLaunchGroupInfo", string(data)}, " ")
}
