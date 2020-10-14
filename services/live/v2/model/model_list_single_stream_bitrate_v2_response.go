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
type ListSingleStreamBitrateV2Response struct {
	// 用量详情。
	BitrateInfoList *[]V2BitrateInfo `json:"bitrate_info_list,omitempty"`
	XRequestId      *string          `json:"X-request-id,omitempty"`
}

func (o ListSingleStreamBitrateV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSingleStreamBitrateV2Response", string(data)}, " ")
}
