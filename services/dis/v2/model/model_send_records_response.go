package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendRecordsResponse Response Object
type SendRecordsResponse struct {

	// 上传失败的数据数量。
	FailedRecordCount *int32 `json:"failed_record_count,omitempty"`

	// 上传结果列表。
	Records        *[]PutRecordsResultEntry `json:"records,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o SendRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendRecordsResponse struct{}"
	}

	return strings.Join([]string{"SendRecordsResponse", string(data)}, " ")
}
