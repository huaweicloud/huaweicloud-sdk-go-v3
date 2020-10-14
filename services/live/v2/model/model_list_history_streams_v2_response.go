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
type ListHistoryStreamsV2Response struct {
	// 历史流信息列表。
	HistoryStreamList *[]HistoryStreamInfo `json:"history_stream_list,omitempty"`
	// 总记录数
	Total      *int32  `json:"total,omitempty"`
	XRequestId *string `json:"X-request-id,omitempty"`
}

func (o ListHistoryStreamsV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListHistoryStreamsV2Response", string(data)}, " ")
}
