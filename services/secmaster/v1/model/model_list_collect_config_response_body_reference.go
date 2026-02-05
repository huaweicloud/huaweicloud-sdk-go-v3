package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectConfigResponseBodyReference 名称和描述
type ListCollectConfigResponseBodyReference struct {

	// 云产品名称
	CsvcDisplay *string `json:"csvc_display,omitempty"`

	// 云产品描述
	CsvcHelp *string `json:"csvc_help,omitempty"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 日志名称
	SourceDisplay *string `json:"source_display,omitempty"`

	// 日志描述
	SourceHelp *string `json:"source_help,omitempty"`
}

func (o ListCollectConfigResponseBodyReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyReference struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyReference", string(data)}, " ")
}
