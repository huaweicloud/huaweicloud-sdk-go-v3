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
type ShowStreamPortraitResponse struct {
	// 播放画像信息列表。
	StreamPortraits *[]StreamPortrait `json:"stream_portraits,omitempty"`
	XRequestId      *string           `json:"X-request-id,omitempty"`
}

func (o ShowStreamPortraitResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowStreamPortraitResponse", string(data)}, " ")
}
