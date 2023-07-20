package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueuePlanResponse Response Object
type CreateQueuePlanResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQueuePlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueuePlanResponse struct{}"
	}

	return strings.Join([]string{"CreateQueuePlanResponse", string(data)}, " ")
}
