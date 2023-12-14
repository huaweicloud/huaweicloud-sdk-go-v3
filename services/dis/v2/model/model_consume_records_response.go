package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsumeRecordsResponse Response Object
type ConsumeRecordsResponse struct {

	// 下载的记录列表。
	Records *[]Record `json:"records,omitempty"`

	// 下一个迭代器。  说明：  数据游标有效期为5分钟。
	NextPartitionCursor *string `json:"next_partition_cursor,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ConsumeRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumeRecordsResponse struct{}"
	}

	return strings.Join([]string{"ConsumeRecordsResponse", string(data)}, " ")
}
