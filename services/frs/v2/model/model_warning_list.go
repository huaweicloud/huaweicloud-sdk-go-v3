package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WarningList struct {
	// 警告ID。

	WarningCode *int32 `json:"warningCode,omitempty"`
	// 告警消息。

	WarningMsg *string `json:"warningMsg,omitempty"`
}

func (o WarningList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarningList struct{}"
	}

	return strings.Join([]string{"WarningList", string(data)}, " ")
}
