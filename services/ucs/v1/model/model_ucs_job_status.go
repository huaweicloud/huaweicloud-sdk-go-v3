package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsJobStatus struct {

	// Job状态： - Running：运行中 - TimedOut：运行超时 - Succeeded：运行成功 - Failed：运行失败
	Status *string `json:"status,omitempty"`

	// 原因
	Reason *string `json:"reason,omitempty"`

	// 开始时间
	StartTime *sdktime.SdkTime `json:"startTime,omitempty"`

	// 结束时间
	FinishTime *sdktime.SdkTime `json:"finishTime,omitempty"`
}

func (o UcsJobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsJobStatus struct{}"
	}

	return strings.Join([]string{"UcsJobStatus", string(data)}, " ")
}
