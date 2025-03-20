package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInterruptRecordsResponse Response Object
type ListInterruptRecordsResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 列表信息
	InterruptRecordList *[]InterruptRecord `json:"interrupt_record_list,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o ListInterruptRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInterruptRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListInterruptRecordsResponse", string(data)}, " ")
}
