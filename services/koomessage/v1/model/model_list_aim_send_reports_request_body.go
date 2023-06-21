package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询智能信息报表请求体。
type ListAimSendReportsRequestBody struct {

	// 报表类型。  - 1：日报表 - 2：月报表  > 若不填，默认是1，即查询日报表。
	ReportType string `json:"report_type"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 开始时间。格式为：2022-05-01T00:00:00Z。
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间。格式为：2022-05-01T00:00:00Z。
	EndTime *string `json:"end_time,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。  > 为提高查询效率，offset+limit须小于等于10000，超出范围查询为空。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAimSendReportsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendReportsRequestBody struct{}"
	}

	return strings.Join([]string{"ListAimSendReportsRequestBody", string(data)}, " ")
}
