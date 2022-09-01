package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowScanJobsRequest struct {

	// 页码
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 内容
	Content *string `json:"content,omitempty" xml:"content"`

	// 预留，待启用
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 预留，待启用
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o ShowScanJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanJobsRequest struct{}"
	}

	return strings.Join([]string{"ShowScanJobsRequest", string(data)}, " ")
}
