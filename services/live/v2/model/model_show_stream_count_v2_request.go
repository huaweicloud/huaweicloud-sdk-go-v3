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
type ShowStreamCountV2Request struct {
	PublishDomains []string `json:"publish_domains"`
	StartTime      *string  `json:"start_time,omitempty"`
	EndTime        *string  `json:"end_time,omitempty"`
}

func (o ShowStreamCountV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowStreamCountV2Request", string(data)}, " ")
}
