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

// Request Object
type ShowStreamPortraitRequest struct {
	PlayDomain string  `json:"play_domain"`
	Stream     *string `json:"stream,omitempty"`
	Time       string  `json:"time"`
}

func (o ShowStreamPortraitRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowStreamPortraitRequest", string(data)}, " ")
}
