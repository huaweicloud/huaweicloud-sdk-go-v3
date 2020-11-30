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
type ListStatisticsResponse struct {
	// 月度调用次数
	Count *[]MonthUsed `json:"count,omitempty"`
	// 月度资源用量
	Gbs            *[]MonthUsed                        `json:"gbs,omitempty"`
	Statistics     *ListFunctionStatisticsResponseBody `json:"statistics,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListStatisticsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListStatisticsResponse", string(data)}, " ")
}
