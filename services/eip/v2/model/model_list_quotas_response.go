/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListQuotasResponse struct {
	Quotas *ResourceResp `json:"quotas,omitempty"`
}

func (o ListQuotasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListQuotasResponse", string(data)}, " ")
}
