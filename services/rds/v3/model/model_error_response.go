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

type ErrorResponse struct {
	// 错误码。
	ErrorCode string `json:"error_code"`
	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o ErrorResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ErrorResponse", string(data)}, " ")
}
