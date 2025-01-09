package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TargetInfo struct {
	TargetType *ScheduleTaskTargetTypeEnum `json:"target_type"`

	// 对象id。
	TargetId string `json:"target_id"`

	// 对象名称。
	TargetName *string `json:"target_name,omitempty"`
}

func (o TargetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetInfo struct{}"
	}

	return strings.Join([]string{"TargetInfo", string(data)}, " ")
}
