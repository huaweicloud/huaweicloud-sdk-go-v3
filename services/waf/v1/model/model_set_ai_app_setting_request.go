package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAiAppSettingRequest Request Object
type SetAiAppSettingRequest struct {
	Body *AiAppSettingRequest `json:"body,omitempty"`
}

func (o SetAiAppSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAiAppSettingRequest struct{}"
	}

	return strings.Join([]string{"SetAiAppSettingRequest", string(data)}, " ")
}
