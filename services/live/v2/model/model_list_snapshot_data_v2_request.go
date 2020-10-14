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
type ListSnapshotDataV2Request struct {
	PublishDomain *string `json:"publish_domain,omitempty"`
	StartTime     *string `json:"start_time,omitempty"`
	EndTime       *string `json:"end_time,omitempty"`
}

func (o ListSnapshotDataV2Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSnapshotDataV2Request", string(data)}, " ")
}
