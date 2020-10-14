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
type ListUsersOfStreamV2Response struct {
	// 域名对应的流量汇总列表。
	DataList   *[]V2UserData `json:"data_list,omitempty"`
	XRequestId *string       `json:"X-request-id,omitempty"`
}

func (o ListUsersOfStreamV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListUsersOfStreamV2Response", string(data)}, " ")
}
