package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InsertQueuePropertyRequestBody 新加队列属性值时对应的属性信息
type InsertQueuePropertyRequestBody struct {
	Properties *InsertQueuePropertyRequestBodyProperties `json:"properties,omitempty"`
}

func (o InsertQueuePropertyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsertQueuePropertyRequestBody struct{}"
	}

	return strings.Join([]string{"InsertQueuePropertyRequestBody", string(data)}, " ")
}
