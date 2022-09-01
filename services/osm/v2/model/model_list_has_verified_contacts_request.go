package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHasVerifiedContactsRequest struct {

	// 联系方式类型
	ContactWay *int32 `json:"contact_way,omitempty" xml:"contact_way"`

	// 客户id
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 子用户id
	SubCustomerId *string `json:"sub_customer_id,omitempty" xml:"sub_customer_id"`

	// 过期时间
	ExpiredTime *string `json:"expired_time,omitempty" xml:"expired_time"`

	// 验证序列号
	VerifiedId *string `json:"verified_id,omitempty" xml:"verified_id"`

	// 联系方式的值
	ContactValue *string `json:"contact_value,omitempty" xml:"contact_value"`

	// 国家码
	AreaCode *string `json:"area_code,omitempty" xml:"area_code"`

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

func (o ListHasVerifiedContactsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHasVerifiedContactsRequest struct{}"
	}

	return strings.Join([]string{"ListHasVerifiedContactsRequest", string(data)}, " ")
}
