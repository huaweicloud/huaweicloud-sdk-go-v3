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
type ListSingleStreamFramerateV2Response struct {
	// 用量详情。
	FramerateInfoList *[]V2FramerateInfo `json:"framerate_info_list,omitempty"`
	XRequestId        *string            `json:"X-request-id,omitempty"`
}

func (o ListSingleStreamFramerateV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSingleStreamFramerateV2Response", string(data)}, " ")
}
