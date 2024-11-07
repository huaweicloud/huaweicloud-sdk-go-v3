package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAsyncHistoryCommandsResponse Response Object
type CountAsyncHistoryCommandsResponse struct {

	// 满足查询条件的记录总数。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountAsyncHistoryCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAsyncHistoryCommandsResponse struct{}"
	}

	return strings.Join([]string{"CountAsyncHistoryCommandsResponse", string(data)}, " ")
}
