/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CancelResourcesSubscriptionResponse struct {
	// |参数名称：退订资源生成的订单ID的列表。| |参数约束以及描述：续订资源生成的订单ID的列表。|
	OrderIds *[]string `json:"order_ids,omitempty"`
}

func (o CancelResourcesSubscriptionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelResourcesSubscriptionResponse", string(data)}, " ")
}
