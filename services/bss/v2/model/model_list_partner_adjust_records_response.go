package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPartnerAdjustRecordsResponse struct {
	// 返回总条数。

	TotalCount *int32 `json:"total_count,omitempty"`
	// 调账记录列表。 具体请参见表2。

	Records        *[]AdjustRecordV2 `json:"records,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListPartnerAdjustRecordsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPartnerAdjustRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPartnerAdjustRecordsResponse", string(data)}, " ")
}
