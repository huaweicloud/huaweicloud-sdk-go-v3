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
type ListRecordDataV2Response struct {
	// 采样数据列表。
	RecordDataList *[]RecordData `json:"record_data_list,omitempty"`
	XRequestId     *string       `json:"X-request-id,omitempty"`
}

func (o ListRecordDataV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRecordDataV2Response", string(data)}, " ")
}
