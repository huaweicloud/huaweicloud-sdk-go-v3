package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWorkitemStatusRecordsV4Response struct {

	// 操作历史
	Records *[]WorkitemStatusRecords `json:"records,omitempty" xml:"records"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWorkitemStatusRecordsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkitemStatusRecordsV4Response struct{}"
	}

	return strings.Join([]string{"ListWorkitemStatusRecordsV4Response", string(data)}, " ")
}
