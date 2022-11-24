package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowScanJobsRequest struct {

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 预留，待启用
	StartTime *string `json:"start_time,omitempty"`

	// 预留，待启用
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowScanJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanJobsRequest struct{}"
	}

	return strings.Join([]string{"ShowScanJobsRequest", string(data)}, " ")
}
