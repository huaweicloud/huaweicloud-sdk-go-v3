/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCollationsResponse struct {
	// 字符集信息列表
	CharSets       *[]string `json:"charSets,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCollationsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCollationsResponse", string(data)}, " ")
}
