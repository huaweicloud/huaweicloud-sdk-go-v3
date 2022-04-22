package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOpLogsResponse struct {

	// 任务列表
	OperationLogs *[]OperationLog `json:"operation_logs,omitempty"`

	// 任务个数
	Count *int32 `json:"count,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOpLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpLogsResponse struct{}"
	}

	return strings.Join([]string{"ListOpLogsResponse", string(data)}, " ")
}
