package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCasesRequest Request Object
type ListCasesRequest struct {

	// 关键字查询，支持多个空格隔开
	SearchKey *[]string `json:"search_key,omitempty"`

	// 标签列表，最多支持5个
	LabelIdList *[]string `json:"label_id_list,omitempty"`

	// app关键字查询
	AppKey *string `json:"app_key,omitempty"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty"`

	// 查询开始时间
	QueryStartTime *string `json:"query_start_time,omitempty"`

	// 查询结束时间
	QueryEndTime *string `json:"query_end_time,omitempty"`

	// 状态 0：待受理 1：处理中 2：待确认结果 3：已完成 4：已撤销 12：无效 17： 待反馈
	Status *int32 `json:"status,omitempty"`

	// 状态列表
	IncidentStatus *string `json:"incident_status,omitempty"`

	// 子用户id
	XCustomerId *string `json:"x_customer_id,omitempty"`

	// 子用户名称
	XCustomerName *string `json:"x_customer_name,omitempty"`

	// 华为云IAM组id，同组其他工单时，该id必传
	GroupId *string `json:"group_id,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 查询限制数量
	Limit *int32 `json:"limit,omitempty"`

	// 对接站点信息。  0（中国站） 1（国际站），不填的话默认为0。
	XSite *int32 `json:"X-Site,omitempty"`

	// 语言环境，值为通用的语言描述字符串，比如zh-cn等，默认为zh-cn。  会根据语言环境对应展示一些国际化的信息，比如工单类型名称等。
	XLanguage *string `json:"X-Language,omitempty"`

	// 环境时区，值为通用的时区描述字符串，比如GMT+8等，默认为GMT+8。  涉及时间的数据会根据环境时区处理。
	XTimeZone *string `json:"X-Time-Zone,omitempty"`
}

func (o ListCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCasesRequest struct{}"
	}

	return strings.Join([]string{"ListCasesRequest", string(data)}, " ")
}
