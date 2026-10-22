package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobUpdateRequestBody 训练作业的描述、事件通知、作业优先级、允许被抢占等信息。
type JobUpdateRequestBody struct {

	// 对训练作业的描述，默认为“NULL”，字符串的长度限制为[0, 256]。
	Description *string `json:"description,omitempty"`

	Spec *JobUpdateRequestBodySpec `json:"spec,omitempty"`

	Endpoints *JobEndpointsReq `json:"endpoints,omitempty"`

	Metadata *JobUpdateRequestBodyMetadata `json:"metadata,omitempty"`
}

func (o JobUpdateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobUpdateRequestBody struct{}"
	}

	return strings.Join([]string{"JobUpdateRequestBody", string(data)}, " ")
}
