package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskResponse Response Object
type UpdateTaskResponse struct {

	// 更新指定ID的迁移任务
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskResponse", string(data)}, " ")
}
