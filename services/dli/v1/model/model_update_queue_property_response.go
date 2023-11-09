package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueuePropertyResponse Response Object
type UpdateQueuePropertyResponse struct {

	// 是否更新成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 接口相关说明
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateQueuePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueuePropertyResponse struct{}"
	}

	return strings.Join([]string{"UpdateQueuePropertyResponse", string(data)}, " ")
}
