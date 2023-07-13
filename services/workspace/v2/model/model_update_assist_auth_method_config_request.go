package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssistAuthMethodConfigRequest Request Object
type UpdateAssistAuthMethodConfigRequest struct {
	Body *AssistAuthMethodConfigRequest `json:"body,omitempty"`
}

func (o UpdateAssistAuthMethodConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssistAuthMethodConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssistAuthMethodConfigRequest", string(data)}, " ")
}
