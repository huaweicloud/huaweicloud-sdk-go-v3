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
type ListPrivateipsResponse struct {
	// 私有IP列表对象
	Privateips []Privateip `json:"privateips,omitempty"`
}

func (o ListPrivateipsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPrivateipsResponse", string(data)}, " ")
}
