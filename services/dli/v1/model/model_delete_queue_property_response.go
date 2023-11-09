package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueuePropertyResponse Response Object
type DeleteQueuePropertyResponse struct {

	// 是否更新成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 接口相关说明
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteQueuePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueuePropertyResponse struct{}"
	}

	return strings.Join([]string{"DeleteQueuePropertyResponse", string(data)}, " ")
}
