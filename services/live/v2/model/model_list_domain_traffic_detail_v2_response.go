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
type ListDomainTrafficDetailV2Response struct {
	// 采样数据列表。
	DataList   *[]TrafficData `json:"data_list,omitempty"`
	XRequestId *string        `json:"X-request-id,omitempty"`
}

func (o ListDomainTrafficDetailV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDomainTrafficDetailV2Response", string(data)}, " ")
}
