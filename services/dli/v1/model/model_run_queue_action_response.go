package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunQueueActionResponse Response Object
type RunQueueActionResponse struct {

	// 请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 当force为true的时候返回的jobId，可以根据job_id的状态来看最终结果
	JobId *string `json:"job_id,omitempty"`

	// 扩缩容的队列名称。
	QueueName *string `json:"queue_name,omitempty"`

	// 扩缩容结果
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o RunQueueActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueueActionResponse struct{}"
	}

	return strings.Join([]string{"RunQueueActionResponse", string(data)}, " ")
}
