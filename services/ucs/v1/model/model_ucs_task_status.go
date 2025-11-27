package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsTaskStatus struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 当前状态原因
	Reason *string `json:"reason,omitempty"`

	// 开始时间
	StartTime *sdktime.SdkTime `json:"startTime,omitempty"`

	// 结束时间
	FinishTime *sdktime.SdkTime `json:"finishTime,omitempty"`
}

func (o UcsTaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsTaskStatus struct{}"
	}

	return strings.Join([]string{"UcsTaskStatus", string(data)}, " ")
}
