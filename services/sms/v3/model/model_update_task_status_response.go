package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskStatusResponse Response Object
type UpdateTaskStatusResponse struct {

	// 管理迁移任务成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResponse", string(data)}, " ")
}
