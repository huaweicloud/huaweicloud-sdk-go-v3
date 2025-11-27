package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInspectionReportResponse Response Object
type ListInspectionReportResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 批量巡检报告列表
	BatchInspectionReportList *[]BatchInspectionReport `json:"batch_inspection_report_list,omitempty"`
	HttpStatusCode            int                      `json:"-"`
}

func (o ListInspectionReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInspectionReportResponse struct{}"
	}

	return strings.Join([]string{"ListInspectionReportResponse", string(data)}, " ")
}
