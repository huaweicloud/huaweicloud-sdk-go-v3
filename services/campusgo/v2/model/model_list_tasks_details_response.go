package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTasksDetailsResponse struct {
	// 符合检索条件的总条目数

	Count *int64 `json:"count,omitempty"`
	// 检索到的服务作业条目

	Tasks          *[]TaskDetails `json:"tasks,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListTasksDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListTasksDetailsResponse", string(data)}, " ")
}
