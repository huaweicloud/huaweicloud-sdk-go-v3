package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSendEmailRequestBody struct {

	// 报告ID列表
	TaskIdList []string `json:"task_id_list"`

	// 邮箱地址
	Email *string `json:"email,omitempty"`

	// 主题名称
	Topic *string `json:"topic,omitempty"`

	// 主题标识符
	TopicUrn *string `json:"topic_urn,omitempty"`

	// OBS桶名
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// 服务地址
	ServiceUri *string `json:"service_uri,omitempty"`
}

func (o BatchSendEmailRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendEmailRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSendEmailRequestBody", string(data)}, " ")
}
