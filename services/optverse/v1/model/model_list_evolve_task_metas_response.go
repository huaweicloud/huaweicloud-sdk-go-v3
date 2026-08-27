package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEvolveTaskMetasResponse Response Object
type ListEvolveTaskMetasResponse struct {

	// 演化任务结构体列表
	Tasks *[]EvolveTaskRsp `json:"tasks,omitempty"`

	// 查询列表总数
	TotalRecords   *int32 `json:"total_records,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEvolveTaskMetasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEvolveTaskMetasResponse struct{}"
	}

	return strings.Join([]string{"ListEvolveTaskMetasResponse", string(data)}, " ")
}
