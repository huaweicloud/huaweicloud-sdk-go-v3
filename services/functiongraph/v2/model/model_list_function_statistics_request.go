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
type ListFunctionStatisticsRequest struct {
	FuncUrn string `json:"func_urn"`
	Period  string `json:"period"`
}

func (o ListFunctionStatisticsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListFunctionStatisticsRequest", string(data)}, " ")
}
