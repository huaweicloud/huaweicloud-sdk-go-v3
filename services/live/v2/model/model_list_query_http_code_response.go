/*
 * LiveAPI
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListQueryHttpCodeResponse struct {
	// 基于时间轴的状态码
	DataSeries *[]HttpCodeSummary `json:"data_series,omitempty"`
	XRequestId *string            `json:"X-request-id,omitempty"`
}

func (o ListQueryHttpCodeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListQueryHttpCodeResponse", string(data)}, " ")
}
