package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MarginInfo 安全报告的页边距
type MarginInfo struct {

	// 安全报告的上边距
	Top *int32 `json:"top,omitempty"`

	// 安全报告的左边距
	Left *int32 `json:"left,omitempty"`

	// 安全报告的下边距
	Bottom *int32 `json:"bottom,omitempty"`

	// 安全报告的右边距
	Right *int32 `json:"right,omitempty"`
}

func (o MarginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MarginInfo struct{}"
	}

	return strings.Join([]string{"MarginInfo", string(data)}, " ")
}
