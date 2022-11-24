package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 购买包周期云模式waf请求参数
type CreatePrepaidCloudWafRequestBody struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 是否自动支付    - false: 否（需要客户手动去支付）   - true：是（自动支付）
	IsAutoPay bool `json:"is_auto_pay"`

	// 是否自动续订   -  true：自动续订   - false：不自动续订
	IsAutoRenew bool `json:"is_auto_renew"`

	// region Id
	RegionId string `json:"region_id"`

	WafProductInfo *WafProductInfo `json:"waf_product_info,omitempty"`

	DomainExpackProductInfo *ExpackProductInfo `json:"domain_expack_product_info,omitempty"`

	BandwidthExpackProductInfo *ExpackProductInfo `json:"bandwidth_expack_product_info,omitempty"`

	RuleExpackProductInfo *ExpackProductInfo `json:"rule_expack_product_info,omitempty"`
}

func (o CreatePrepaidCloudWafRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrepaidCloudWafRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrepaidCloudWafRequestBody", string(data)}, " ")
}
