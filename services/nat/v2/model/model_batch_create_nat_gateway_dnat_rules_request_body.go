package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateNatGatewayDnatRulesRequestBody This is an auto create Body Object
type BatchCreateNatGatewayDnatRulesRequestBody struct {

	// DNAT规则批量创建对象的请求体。
	DnatRules []CreateNatGatewayDnatOption `json:"dnat_rules"`
}

func (o BatchCreateNatGatewayDnatRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNatGatewayDnatRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesRequestBody", string(data)}, " ")
}
