package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFutureExecutionsResponse Response Object
type ListFutureExecutionsResponse struct {

	// 未来执行的具体时间列表。
	FutureExecutions *[]string `json:"future_executions,omitempty"`
	HttpStatusCode   int       `json:"-"`
}

func (o ListFutureExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFutureExecutionsResponse struct{}"
	}

	return strings.Join([]string{"ListFutureExecutionsResponse", string(data)}, " ")
}
