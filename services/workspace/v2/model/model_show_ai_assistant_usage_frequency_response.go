package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiAssistantUsageFrequencyResponse Response Object
type ShowAiAssistantUsageFrequencyResponse struct {

	// 用户使用频次统计列表。
	Users          *[]UserUsageFrequencyInfo `json:"users,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowAiAssistantUsageFrequencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiAssistantUsageFrequencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAiAssistantUsageFrequencyResponse", string(data)}, " ")
}
