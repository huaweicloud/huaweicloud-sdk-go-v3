package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCollectConfigResponseBodySourceList struct {

	// 云产品名称
	CsvcDisplay *string `json:"csvc_display,omitempty"`

	// 云产品描述
	CsvcHzzelp *string `json:"csvc_hzzelp,omitempty"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 日志名称
	SourceDisplay *string `json:"source_display,omitempty"`

	// 日志描述
	SourceHelp *string `json:"source_help,omitempty"`
}

func (o ListCollectConfigResponseBodySourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodySourceList struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodySourceList", string(data)}, " ")
}
