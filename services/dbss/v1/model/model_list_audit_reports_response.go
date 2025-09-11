package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditReportsResponse Response Object
type ListAuditReportsResponse struct {

	// 报表对象列表
	Reports *[]ReportInfo `json:"reports,omitempty"`

	// 总记录数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditReportsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditReportsResponse", string(data)}, " ")
}
