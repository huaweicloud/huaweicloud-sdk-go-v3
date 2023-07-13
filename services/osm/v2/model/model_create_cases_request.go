package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCasesRequest Request Object
type CreateCasesRequest struct {

	// 对接站点信息。  0（中国站） 1（国际站），不填的话默认为0。
	XSite *int32 `json:"X-Site,omitempty"`

	// 语言环境，值为通用的语言描述字符串，比如zh-cn等，默认为zh-cn。  会根据语言环境对应展示一些国际化的信息，比如工单类型名称等。
	XLanguage *string `json:"X-Language,omitempty"`

	// 环境时区，值为通用的时区描述字符串，比如GMT+8等，默认为GMT+8。  涉及时间的数据会根据环境时区处理。
	XTimeZone *string `json:"X-Time-Zone,omitempty"`

	// 手机验证序列号id，如果是非注册联系方式则必填，通过\"查询已验证的列表\"接口查询获取
	XPhoneVerifiedid *string `json:"x-phone-verifiedid,omitempty"`

	// 邮件验证序列号id，如果是非注册联系方式则必填，通过\"查询已验证的列表\"接口查询获取
	XEmailVerifiedid *string `json:"x-email-verifiedid,omitempty"`

	Body *CreateOrderIncidentV2Req `json:"body,omitempty"`
}

func (o CreateCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCasesRequest struct{}"
	}

	return strings.Join([]string{"CreateCasesRequest", string(data)}, " ")
}
