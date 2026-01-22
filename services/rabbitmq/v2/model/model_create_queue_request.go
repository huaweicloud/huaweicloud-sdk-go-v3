package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueueRequest Request Object
type CreateQueueRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// vhost名称，名称中包含/时，需要将/替换为__F_SLASH__，否则会调用失败。例如：Vhost名称为/test，入参值为__F_SLASH__test。
	Vhost string `json:"vhost"`

	Body *CreateQueueBody `json:"body,omitempty"`
}

func (o CreateQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueRequest struct{}"
	}

	return strings.Join([]string{"CreateQueueRequest", string(data)}, " ")
}
