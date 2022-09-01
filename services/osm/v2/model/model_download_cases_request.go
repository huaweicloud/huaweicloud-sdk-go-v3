package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadCasesRequest struct {

	// 语言
	Language string `json:"language" xml:"language"`

	// 时区
	Timezone string `json:"timezone" xml:"timezone"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty" xml:"incident_id"`

	// 查询开始时间
	QueryStartTime *string `json:"query_start_time,omitempty" xml:"query_start_time"`

	// 查询结束时间
	QueryEndTime *string `json:"query_end_time,omitempty" xml:"query_end_time"`

	// 子用户名称
	XCustomerName *string `json:"x_customer_name,omitempty" xml:"x_customer_name"`

	// 搜索关键字
	SearchKey *string `json:"search_key,omitempty" xml:"search_key"`

	// 状态 0：待受理 1：处理中 2：待确认结果 3：已完成 4：已撤销 12：无效 17： 待反馈
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 用户id
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 来源id
	TenantSourceIdList *[]string `json:"tenant_source_id_list,omitempty" xml:"tenant_source_id_list"`

	// 子用户id
	SubCustomerId *string `json:"sub_customer_id,omitempty" xml:"sub_customer_id"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 对接站点信息。  0（中国站） 1（国际站），不填的话默认为0。
	XSite *int32 `json:"X-Site,omitempty" xml:"X-Site"`

	// 语言环境，值为通用的语言描述字符串，比如zh-cn等，默认为zh-cn。  会根据语言环境对应展示一些国际化的信息，比如工单类型名称等。
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 环境时区，值为通用的时区描述字符串，比如GMT+8等，默认为GMT+8。  涉及时间的数据会根据环境时区处理。
	XTimeZone *string `json:"X-Time-Zone,omitempty" xml:"X-Time-Zone"`
}

func (o DownloadCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadCasesRequest struct{}"
	}

	return strings.Join([]string{"DownloadCasesRequest", string(data)}, " ")
}
