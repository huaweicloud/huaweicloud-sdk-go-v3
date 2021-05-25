package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListQuotasRequest struct {
	// 功能说明：根据type过滤查询指定类型的配额 取值范围：vpc，subnet，securityGroup，securityGroupRule，publicIp，vpn，vpngw，vpcPeer，firewall，shareBandwidth，shareBandwidthIP，loadbalancer，listener

	Type *string `json:"type,omitempty"`
}

func (o ListQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}
