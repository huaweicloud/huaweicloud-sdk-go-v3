package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GetCursorResponse struct {
	// 数据游标。  取值范围：1~512个字符。  说明：  数据游标有效期为5分钟。

	PartitionCursor *string `json:"partition_cursor,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o GetCursorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCursorResponse struct{}"
	}

	return strings.Join([]string{"GetCursorResponse", string(data)}, " ")
}
