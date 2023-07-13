package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePrepaidCloudWafRequestBody 变更包周期云模式waf规格请求体
type ChangePrepaidCloudWafRequestBody struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 是否自动支付    - false: 否（需要客户手动去支付）    - true：是（自动支付）
	IsAutoPay bool `json:"is_auto_pay"`

	WafProductInfo *AlterWafProductInfo `json:"waf_product_info,omitempty"`

	DomainExpackProductInfo *ExpackProductInfo `json:"domain_expack_product_info,omitempty"`

	BandwidthExpackProductInfo *ExpackProductInfo `json:"bandwidth_expack_product_info,omitempty"`

	RuleExpackProductInfo *ExpackProductInfo `json:"rule_expack_product_info,omitempty"`
}

func (o ChangePrepaidCloudWafRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePrepaidCloudWafRequestBody struct{}"
	}

	return strings.Join([]string{"ChangePrepaidCloudWafRequestBody", string(data)}, " ")
}
