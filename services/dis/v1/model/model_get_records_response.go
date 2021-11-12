package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GetRecordsResponse struct {
	Records *[]Record `json:"records,omitempty"`
	// 下一个迭代器。  说明：  数据游标有效期为5分钟。

	NextPartitionCursor *string `json:"next_partition_cursor,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o GetRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetRecordsResponse struct{}"
	}

	return strings.Join([]string{"GetRecordsResponse", string(data)}, " ")
}
