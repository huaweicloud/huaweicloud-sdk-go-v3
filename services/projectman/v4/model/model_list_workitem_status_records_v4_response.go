package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkitemStatusRecordsV4Response Response Object
type ListWorkitemStatusRecordsV4Response struct {

	// 操作历史
	Records *[]WorkitemStatusRecords `json:"records,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWorkitemStatusRecordsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkitemStatusRecordsV4Response struct{}"
	}

	return strings.Join([]string{"ListWorkitemStatusRecordsV4Response", string(data)}, " ")
}
