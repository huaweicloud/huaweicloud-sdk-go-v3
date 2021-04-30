package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListOpRecordResponse struct {
	// 操作记录总数

	Count *int64 `json:"count,omitempty"`
	// 操作记录列表

	OperationRecords *[]RecordDetailInfo `json:"operation_records,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ListOpRecordResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOpRecordResponse struct{}"
	}

	return strings.Join([]string{"ListOpRecordResponse", string(data)}, " ")
}
