package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportsResponse Response Object
type ListReportsResponse struct {

	// 本次分页查询响应的报告条数
	Count *int64 `json:"count,omitempty"`

	// 下一页的marker
	NextMarker *string `json:"next_marker,omitempty"`

	// 报告列表
	Reports        *[]ReportEntity `json:"reports,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportsResponse struct{}"
	}

	return strings.Join([]string{"ListReportsResponse", string(data)}, " ")
}
