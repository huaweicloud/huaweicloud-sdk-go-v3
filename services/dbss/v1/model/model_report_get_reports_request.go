package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportGetReportsRequest struct {

	// 报表名称
	Name *string `json:"name,omitempty"`

	// 当前页码
	Page int32 `json:"page"`

	// 每页条数
	Size int32 `json:"size"`
}

func (o ReportGetReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportGetReportsRequest struct{}"
	}

	return strings.Join([]string{"ReportGetReportsRequest", string(data)}, " ")
}
