package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskSpeedResponse Response Object
type UpdateTaskSpeedResponse struct {

	// 上报数据迁移进度和速率成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTaskSpeedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSpeedResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskSpeedResponse", string(data)}, " ")
}
