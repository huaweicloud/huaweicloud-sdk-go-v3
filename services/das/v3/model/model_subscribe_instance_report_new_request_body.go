package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeInstanceReportNewRequestBody 订阅实例报告请求体
type SubscribeInstanceReportNewRequestBody struct {

	// 协议
	Protocol string `json:"protocol"`

	// 地址
	Endpoint string `json:"endpoint"`

	// 主题
	Topic string `json:"topic"`

	// 主题地址
	TopicUrn string `json:"topic_urn"`

	// 桶名
	BucketName *string `json:"bucket_name,omitempty"`

	// 风险等级
	Level string `json:"level"`

	// 服务URI
	ServiceUri string `json:"service_uri"`
}

func (o SubscribeInstanceReportNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeInstanceReportNewRequestBody struct{}"
	}

	return strings.Join([]string{"SubscribeInstanceReportNewRequestBody", string(data)}, " ")
}
