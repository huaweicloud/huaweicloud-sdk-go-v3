package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssistAuthConfigApplyObjectsRequest Request Object
type UpdateAssistAuthConfigApplyObjectsRequest struct {
	Body *AssistAuthApplyObjectsRequest `json:"body,omitempty"`
}

func (o UpdateAssistAuthConfigApplyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssistAuthConfigApplyObjectsRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssistAuthConfigApplyObjectsRequest", string(data)}, " ")
}
