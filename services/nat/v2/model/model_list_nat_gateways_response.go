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

// Response Object
type ListNatGatewaysResponse struct {
	// 查询公网NAT网关实例列表的响应体。 详见NatGateway字段说明。
	NatGateways    *[]NatGatewayResponseBody `json:"nat_gateways,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListNatGatewaysResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListNatGatewaysResponse", string(data)}, " ")
}
