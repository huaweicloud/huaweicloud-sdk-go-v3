package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTopicOrBatchDeleteTopicRequest Request Object
type CreateTopicOrBatchDeleteTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 批量删除topic时使用，不配置则为创建接口。
	Action *string `json:"action,omitempty"`

	Body *CreateTopicOrBatchDeleteTopicReq `json:"body,omitempty"`
}

func (o CreateTopicOrBatchDeleteTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicOrBatchDeleteTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateTopicOrBatchDeleteTopicRequest", string(data)}, " ")
}
