package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopicConfig 话题控制规则配置
type TopicConfig struct {

	// 话题名称
	TopicName string `json:"topic_name"`

	// 话题描述
	TopicDescription string `json:"topic_description"`

	// 敏感词列表
	SensitiveWords *[]string `json:"sensitive_words,omitempty"`

	// 黑样本列表
	BlackSamples *[]string `json:"black_samples,omitempty"`

	// 白样本列表
	WhiteSamples *[]string `json:"white_samples,omitempty"`

	// 检测作用域，取值 request/response/request_response（值合法性由后端validateTopicConfig校验，契约层不声明enum）
	Scope string `json:"scope"`

	// 拦截回包内容（仅外层action.category=block时有意义，非block落库归一化为null）
	BlockResponse *string `json:"block_response,omitempty"`

	// 关联的自定义大模型ID（可选）
	CustomLlmId *string `json:"custom_llm_id,omitempty"`
}

func (o TopicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicConfig struct{}"
	}

	return strings.Join([]string{"TopicConfig", string(data)}, " ")
}
