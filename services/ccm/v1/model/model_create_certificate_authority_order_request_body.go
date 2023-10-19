package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCertificateAuthorityOrderRequestBody struct {

	// 云服务类型，固定为'hws.service.type.ccm'
	CloudServiceType string `json:"cloud_service_type"`

	// 计费模式 0包周期
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期 2月 3年
	PeriodType int32 `json:"period_type"`

	// 订购周期数
	PeriodNum int32 `json:"period_num"`

	// 是否自动续费 1是 0否
	IsAutoRenew int32 `json:"is_auto_renew"`

	// 折扣信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 订购数量
	SubscriptionNum int32 `json:"subscription_num"`

	// 是否自动支付 1是 0否 不填默认为0
	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 产品列表，详情请参见**ProductInfo**字段数据结构说明。
	ProductInfos []ProductInfo `json:"product_infos"`
}

func (o CreateCertificateAuthorityOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateAuthorityOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateAuthorityOrderRequestBody", string(data)}, " ")
}
