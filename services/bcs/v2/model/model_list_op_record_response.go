package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOpRecordResponse struct {

	// 操作记录总数
	Count *int64 `json:"count,omitempty" xml:"count"`

	// 操作记录列表
	OperationRecords *[]RecordDetailInfo `json:"operation_records,omitempty" xml:"operation_records"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ListOpRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpRecordResponse struct{}"
	}

	return strings.Join([]string{"ListOpRecordResponse", string(data)}, " ")
}
