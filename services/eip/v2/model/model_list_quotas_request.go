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

// Request Object
type ListQuotasRequest struct {
	Type *string `json:"type,omitempty"`
}

func (o ListQuotasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}
