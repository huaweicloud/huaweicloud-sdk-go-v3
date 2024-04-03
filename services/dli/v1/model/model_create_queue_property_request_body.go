package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueuePropertyRequestBody 新加队列属性值时对应的属性信息
type CreateQueuePropertyRequestBody struct {
	Properties *CommonQueueProperty `json:"properties,omitempty"`
}

func (o CreateQueuePropertyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueuePropertyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateQueuePropertyRequestBody", string(data)}, " ")
}
