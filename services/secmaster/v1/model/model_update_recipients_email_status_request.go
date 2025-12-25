package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecipientsEmailStatusRequest Request Object
type UpdateRecipientsEmailStatusRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateRecipientsStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateRecipientsEmailStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecipientsEmailStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecipientsEmailStatusRequest", string(data)}, " ")
}
