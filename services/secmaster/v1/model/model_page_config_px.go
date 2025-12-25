package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageConfigPx 安全报告的页边距
type PageConfigPx struct {

	// 安全报告的宽度
	Width *int32 `json:"width,omitempty"`

	// 安全报告的高度
	Height *int32 `json:"height,omitempty"`

	Margin *MarginInfo `json:"margin,omitempty"`
}

func (o PageConfigPx) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageConfigPx struct{}"
	}

	return strings.Join([]string{"PageConfigPx", string(data)}, " ")
}
