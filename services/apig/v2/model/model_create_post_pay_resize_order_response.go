package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPayResizeOrderResponse Response Object
type CreatePostPayResizeOrderResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例扩容任务信息
	Message *string `json:"message,omitempty"`

	// 任务编号
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePostPayResizeOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPayResizeOrderResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPayResizeOrderResponse", string(data)}, " ")
}
