/*
 * NAT
 *
 * Open Api of Public Nat.
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type BatchCreateDnatRulesRequestBody struct {
	// DNAT规则批量创建对象的请求体。
	DnatRules []CreateNatGatewayDnatOption `json:"dnat_rules"`
}

func (o BatchCreateDnatRulesRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateDnatRulesRequestBody", string(data)}, " ")
}
