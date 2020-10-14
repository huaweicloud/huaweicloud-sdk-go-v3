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
type ListRecordDataV2Request struct {
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
}

func (o ListRecordDataV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRecordDataV2Request", string(data)}, " ")
}
