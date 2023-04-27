package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsTemplateResp struct {

	// 模板主键ID，用于获取、修改、删除模板以及查询模板变量的唯一标识
	Id *string `json:"id,omitempty"`

	// 创建时间[yyyy-MM-dd HH:mm:ss]
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间[yyyy-MM-dd HH:mm:ss]
	UpdateTime *string `json:"update_time,omitempty"`

	// 租户customer id
	CustomerId *string `json:"customer_id,omitempty"`

	Tenant *TenantBasic `json:"tenant,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板类型
	TemplateType *string `json:"template_type,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 应用key
	AppKey *string `json:"app_key,omitempty"`

	// 签名主键id
	SignId *string `json:"sign_id,omitempty"`

	// 模板内容
	TemplateContent *string `json:"template_content,omitempty"`

	// 是否有变量
	HasVariable *string `json:"has_variable,omitempty"`

	// 申请描述
	TemplateDesc *string `json:"template_desc,omitempty"`

	// 审核意见
	ReviewDesc *string `json:"review_desc,omitempty"`

	// 审核人账号
	ReviewOrder *string `json:"review_order,omitempty"`

	// 流程状态
	FlowStatus *string `json:"flow_status,omitempty"`

	// 是否是通用模板
	UniversalTemplate *int32 `json:"universal_template,omitempty"`

	// 模板状态
	Status *string `json:"status,omitempty"`

	// 地域 1. cn：国内 2. intl：国际
	Region *string `json:"region,omitempty"`

	// 中括号类型 支持枚举值：  CN: 中文类型  GB: 英文类型
	Brackets *string `json:"brackets,omitempty"`

	// 站点
	Site *string `json:"site,omitempty"`

	// 催审状态
	UrgeStatus *string `json:"urge_status,omitempty"`

	// 催审时间
	UrgeTime *string `json:"urge_time,omitempty"`

	// 催审描述
	UrgeDesc *string `json:"urge_desc,omitempty"`

	// 发送国家1
	SendCountry1 *int64 `json:"send_country1,omitempty"`

	// 发送国家2
	SendCountry2 *int64 `json:"send_country2,omitempty"`

	// 发送国家3
	SendCountry3 *int64 `json:"send_country3,omitempty"`

	// 是否支持多OMP
	IsSupportMultiomp *bool `json:"is_support_multiomp,omitempty"`

	// 国家名称列表，返回发送国家前三名的国家名称，国家名称间以\"~\"分隔
	CountryName *string `json:"country_name,omitempty"`
}

func (o SmsTemplateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsTemplateResp struct{}"
	}

	return strings.Join([]string{"SmsTemplateResp", string(data)}, " ")
}
