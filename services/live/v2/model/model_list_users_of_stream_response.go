/*
 * Live
 *
 * 数据分析服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListUsersOfStreamResponse struct {
	// 域名对应的流量汇总列表。
	DataList       *[]V2UserData `json:"data_list,omitempty"`
	XRequestId     *string       `json:"X-request-id,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListUsersOfStreamResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListUsersOfStreamResponse", string(data)}, " ")
}
