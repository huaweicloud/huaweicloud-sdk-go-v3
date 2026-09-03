package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksByTaskIdResponse Response Object
type ListTasksByTaskIdResponse struct {

	// SQL洞察任务明细
	Tasks *[]SqlParseTask `json:"tasks,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTasksByTaskIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksByTaskIdResponse struct{}"
	}

	return strings.Join([]string{"ListTasksByTaskIdResponse", string(data)}, " ")
}
