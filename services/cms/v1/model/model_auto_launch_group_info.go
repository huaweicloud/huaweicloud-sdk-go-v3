package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 智能购买组信息
type AutoLaunchGroupInfo struct {

	// 智能购买组id
	Id string `json:"id"`

	// autoLaunchGroup的名称(1-64个字符)，只能包含中文、字母、数字、下划线或中划线
	Name string `json:"name"`

	// 请求类型
	Type string `json:"type"`

	// autoLaunchGroup的运行状态 SUBMITTED|ACTIVE|DELETING|DELETED
	Status string `json:"status"`

	// autoLaunchGroup的任务状态， INIT|HANDLING|FULFILLED|ERROR
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
