package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteJobActionResponse struct {

	// 异步操作任务响应查询ID。
	QueryId        *string `json:"query_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteJobActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteJobActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteJobActionResponse", string(data)}, " ")
}
