/*
 * DDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ResizeInstanceRequestBody struct {
	Resize *ResizeInstanceOption `json:"resize"`
}

func (o ResizeInstanceRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResizeInstanceRequestBody", string(data)}, " ")
}
