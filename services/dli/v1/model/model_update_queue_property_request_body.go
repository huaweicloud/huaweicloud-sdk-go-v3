package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueuePropertyRequestBody 队列待更新属性
type UpdateQueuePropertyRequestBody struct {
	Properties *UpdateQueuePropertyRequestBodyProperties `json:"properties,omitempty"`
}

func (o UpdateQueuePropertyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueuePropertyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateQueuePropertyRequestBody", string(data)}, " ")
}
