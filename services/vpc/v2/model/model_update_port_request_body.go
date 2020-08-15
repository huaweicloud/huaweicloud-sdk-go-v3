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
type UpdatePortRequestBody struct {
	Port *UpdatePortOption `json:"port"`
}

func (o UpdatePortRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePortRequestBody", string(data)}, " ")
}
