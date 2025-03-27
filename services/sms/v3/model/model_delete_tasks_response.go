package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTasksResponse Response Object
type DeleteTasksResponse struct {

	// 批量删除迁移任务成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTasksResponse struct{}"
	}

	return strings.Join([]string{"DeleteTasksResponse", string(data)}, " ")
}
