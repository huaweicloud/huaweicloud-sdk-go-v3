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
type ListHistoryStreamsV2Request struct {
	Domain string  `json:"domain"`
	App    *string `json:"app,omitempty"`
	Offset *int32  `json:"offset,omitempty"`
	Limit  *int32  `json:"limit,omitempty"`
}

func (o ListHistoryStreamsV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListHistoryStreamsV2Request", string(data)}, " ")
}
