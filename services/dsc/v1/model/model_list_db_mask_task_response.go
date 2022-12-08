package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDbMaskTaskResponse struct {

	// 脱敏任务列表
	Tasks *[]DbMaskTaskInfo `json:"tasks,omitempty"`

	// 脱敏任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbMaskTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbMaskTaskResponse struct{}"
	}

	return strings.Join([]string{"ListDbMaskTaskResponse", string(data)}, " ")
}
