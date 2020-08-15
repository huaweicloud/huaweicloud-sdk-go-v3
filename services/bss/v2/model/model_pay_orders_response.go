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
type PayOrdersResponse struct {
}

func (o PayOrdersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PayOrdersResponse", string(data)}, " ")
}
