package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTakeOverTaskResponse struct {

	// 托管任务信息。
	Tasks *[]TakeOverTask `json:"tasks,omitempty" xml:"tasks"`

	// 任务数量。
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTakeOverTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTakeOverTaskResponse struct{}"
	}

	return strings.Join([]string{"ListTakeOverTaskResponse", string(data)}, " ")
}
