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

// Response Object
type ListQuotasResponse struct {
	Quotas         *ListQuotasResult `json:"quotas,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListQuotasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListQuotasResponse", string(data)}, " ")
}
