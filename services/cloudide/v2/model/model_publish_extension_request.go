package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishExtensionRequest Request Object
type PublishExtensionRequest struct {

	// 发布商凭证,x-publisher-token和X-Auth-Token必传一个
	XPublisherToken *string `json:"x-publisher-token,omitempty"`

	// 任务id
	TaskId string `json:"task_id"`

	Body *TaskModelMarketPlace `json:"body,omitempty"`
}

func (o PublishExtensionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishExtensionRequest struct{}"
	}

	return strings.Join([]string{"PublishExtensionRequest", string(data)}, " ")
}
