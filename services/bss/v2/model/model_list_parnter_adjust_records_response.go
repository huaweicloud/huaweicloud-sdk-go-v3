package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListParnterAdjustRecordsResponse struct {
	// 调账记录列表。 具体请参见表2。

	Records *[]AdjustRecordV3 `json:"records,omitempty"`
	// 返回总条数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListParnterAdjustRecordsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListParnterAdjustRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListParnterAdjustRecordsResponse", string(data)}, " ")
}
