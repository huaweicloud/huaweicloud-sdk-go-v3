package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WarningList struct {

	// 警告ID。
	WarningCode *int32 `json:"warningCode,omitempty" xml:"warningCode"`

	// 告警消息。
	WarningMsg *string `json:"warningMsg,omitempty" xml:"warningMsg"`
}

func (o WarningList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarningList struct{}"
	}

	return strings.Join([]string{"WarningList", string(data)}, " ")
}
