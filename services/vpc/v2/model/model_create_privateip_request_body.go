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

//
type CreatePrivateipRequestBody struct {
	// 私有IP列表对象
	Privateips []CreatePrivateipOption `json:"privateips"`
}

func (o CreatePrivateipRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePrivateipRequestBody", string(data)}, " ")
}
