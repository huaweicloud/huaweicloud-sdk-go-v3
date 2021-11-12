package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRecordDataResponse struct {
	// 采样数据列表。

	RecordDataList *[]RecordData `json:"record_data_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecordDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordDataResponse struct{}"
	}

	return strings.Join([]string{"ListRecordDataResponse", string(data)}, " ")
}
