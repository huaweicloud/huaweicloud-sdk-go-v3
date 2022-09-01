package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateResOnlineInstanceRequestBody struct {

	// 作业名称，1-64位的字母、数字、下划线、中划线组合。
	JobName string `json:"job_name" xml:"job_name"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 类别: - SERVICE，在线服务
	Category string `json:"category" xml:"category"`

	// 作业类型： - infer，推理服务
	JobType string `json:"job_type" xml:"job_type"`

	JobConfig *JobConfig `json:"job_config" xml:"job_config"`

	// 通知消息配置。
	TopicUrn *string `json:"topicUrn,omitempty" xml:"topicUrn"`
}

func (o CreateResOnlineInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResOnlineInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResOnlineInstanceRequestBody", string(data)}, " ")
}
