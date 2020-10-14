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
type ListSingleStreamFramerateV2Request struct {
	Domain    string  `json:"domain"`
	App       string  `json:"app"`
	Stream    string  `json:"stream"`
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
}

func (o ListSingleStreamFramerateV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSingleStreamFramerateV2Request", string(data)}, " ")
}
