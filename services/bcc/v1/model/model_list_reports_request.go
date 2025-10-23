package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportsRequest Request Object
type ListReportsRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 报告配置id
	ReportSettingId *string `json:"report_setting_id,omitempty"`

	// 生成报告的起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 生成报告的截止时间，例如：2025-03-21T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。
	Marker *string `json:"marker,omitempty"`

	// 单次查询的条数限制,取值范围(0,100]，默认值为10，用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportsRequest struct{}"
	}

	return strings.Join([]string{"ListReportsRequest", string(data)}, " ")
}
