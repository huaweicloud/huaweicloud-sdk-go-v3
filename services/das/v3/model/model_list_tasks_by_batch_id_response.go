package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksByBatchIdResponse Response Object
type ListTasksByBatchIdResponse struct {

	// SQL洞察任务明细
	Tasks *[]SqlParseTask `json:"tasks,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTasksByBatchIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksByBatchIdResponse struct{}"
	}

	return strings.Join([]string{"ListTasksByBatchIdResponse", string(data)}, " ")
}
