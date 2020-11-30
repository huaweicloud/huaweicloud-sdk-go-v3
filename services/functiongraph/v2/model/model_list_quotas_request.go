/*
 * FunctionGraph
 *
 * API v2
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListQuotasRequest struct {
}

func (o ListQuotasRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}
