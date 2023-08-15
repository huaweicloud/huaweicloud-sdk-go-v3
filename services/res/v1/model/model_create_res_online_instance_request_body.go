package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResOnlineInstanceRequestBody This is a auto create Body Object
type CreateResOnlineInstanceRequestBody struct {

	// 作业名称，1-64位的字母、数字、下划线、中划线组合。
	JobName string `json:"job_name"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 类别: - SERVICE，在线服务
	Category string `json:"category"`

	// 作业类型： - infer，推理服务
	JobType string `json:"job_type"`

	JobConfig *JobConfig `json:"job_config"`

	// 通知消息配置。
	TopicUrn *string `json:"topicUrn,omitempty"`
}

func (o CreateResOnlineInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResOnlineInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResOnlineInstanceRequestBody", string(data)}, " ")
}
