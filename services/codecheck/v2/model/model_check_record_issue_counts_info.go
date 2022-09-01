package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckRecordIssueCountsInfo struct {

	// 致命问题
	Critical *int32 `json:"critical,omitempty" xml:"critical"`

	// 严重问题
	Serious *int32 `json:"serious,omitempty" xml:"serious"`

	// 常规问题
	Normal *int32 `json:"normal,omitempty" xml:"normal"`

	// 提示问题
	Prompt *int32 `json:"prompt,omitempty" xml:"prompt"`
}

func (o CheckRecordIssueCountsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRecordIssueCountsInfo struct{}"
	}

	return strings.Join([]string{"CheckRecordIssueCountsInfo", string(data)}, " ")
}
