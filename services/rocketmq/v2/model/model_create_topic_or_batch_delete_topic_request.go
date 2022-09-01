package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTopicOrBatchDeleteTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 批量删除topic时使用，不配置则为创建接口。
	Action *string `json:"action,omitempty" xml:"action"`

	Body *CreateTopicOrBatchDeleteTopicReq `json:"body,omitempty" xml:"body"`
}

func (o CreateTopicOrBatchDeleteTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicOrBatchDeleteTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateTopicOrBatchDeleteTopicRequest", string(data)}, " ")
}
