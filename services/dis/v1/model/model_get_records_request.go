package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GetRecordsRequest struct {
	// 数据游标，需要先通过获取数据游标的接口获取。  取值范围：1~512个字符。  说明：  数据游标有效期为5分钟。

	PartitionCursor string `json:"partition-cursor"`
	// 每个请求获取记录的最大字节数。  注意：  该值如果小于分区中单条记录的大小，会导致一直无法获取到记录。

	MaxFetchBytes *int32 `json:"max_fetch_bytes,omitempty"`
}

func (o GetRecordsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetRecordsRequest struct{}"
	}

	return strings.Join([]string{"GetRecordsRequest", string(data)}, " ")
}
