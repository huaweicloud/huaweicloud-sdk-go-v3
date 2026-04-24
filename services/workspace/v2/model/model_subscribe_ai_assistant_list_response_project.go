package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeAiAssistantListResponseProject 订阅项目信息。
type SubscribeAiAssistantListResponseProject struct {

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// ai 功能是否启用。 * true： 启用 * false： 不启用
	AiFunc *bool `json:"ai_func,omitempty"`
}

func (o SubscribeAiAssistantListResponseProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeAiAssistantListResponseProject struct{}"
	}

	return strings.Join([]string{"SubscribeAiAssistantListResponseProject", string(data)}, " ")
}
