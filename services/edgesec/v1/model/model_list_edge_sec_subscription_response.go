package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeSecSubscriptionResponse Response Object
type ListEdgeSecSubscriptionResponse struct {

	// 已经添加的WAF防护域名数量
	WafDomainNum *int32 `json:"waf_domain_num,omitempty"`

	// 已经添加的WAF IP黑白规则数量
	WafRuleNum *int32 `json:"waf_rule_num,omitempty"`

	// 已经添加的DDoS防护域名数量
	DdosDomainNum *int32 `json:"ddos_domain_num,omitempty"`

	// 产品信息
	ProductInfos   *[]EdgeSecProductResource `json:"product_infos,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListEdgeSecSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeSecSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeSecSubscriptionResponse", string(data)}, " ")
}
