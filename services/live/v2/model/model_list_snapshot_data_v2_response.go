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
type ListSnapshotDataV2Response struct {
	// 采样数据列表。
	SnapshotList *[]SnapshotData `json:"snapshot_list,omitempty"`
	// 指定时间区间内截图数量总和
	Total      *int64  `json:"total,omitempty"`
	XRequestId *string `json:"X-request-id,omitempty"`
}

func (o ListSnapshotDataV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSnapshotDataV2Response", string(data)}, " ")
}
