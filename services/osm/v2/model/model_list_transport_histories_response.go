package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTransportHistoriesResponse struct {
	// 总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 文件记录列表

	FileOpsList    *[]FileOperateLog `json:"file_ops_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTransportHistoriesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTransportHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListTransportHistoriesResponse", string(data)}, " ")
}
