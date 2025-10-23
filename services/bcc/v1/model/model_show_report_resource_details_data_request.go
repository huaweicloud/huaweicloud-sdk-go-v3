package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportResourceDetailsDataRequest Request Object
type ShowReportResourceDetailsDataRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 报告ID
	ReportId string `json:"report_id"`

	// 单次查询的条数限制,取值范围(0,100]，默认值为10，用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`

	// 请求游标
	Marker *string `json:"marker,omitempty"`
}

func (o ShowReportResourceDetailsDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportResourceDetailsDataRequest struct{}"
	}

	return strings.Join([]string{"ShowReportResourceDetailsDataRequest", string(data)}, " ")
}
