/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type CreatePrivateipResponse struct {
	// 私有IP列表对象
	Privateips []Privateip `json:"privateips,omitempty"`
}

func (o CreatePrivateipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePrivateipResponse", string(data)}, " ")
}
