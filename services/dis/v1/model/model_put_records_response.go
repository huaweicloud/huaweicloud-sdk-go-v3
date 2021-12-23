package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PutRecordsResponse struct {
	// 上传失败的数据数量。

	FailedRecordCount *int32 `json:"failed_record_count,omitempty"`

	Records        *[]PutRecordsResultEntry `json:"records,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o PutRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutRecordsResponse struct{}"
	}

	return strings.Join([]string{"PutRecordsResponse", string(data)}, " ")
}
