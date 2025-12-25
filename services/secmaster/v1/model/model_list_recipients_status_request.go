package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecipientsStatusRequest Request Object
type ListRecipientsStatusRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *ListRecipientsStatusRequestBody `json:"body,omitempty"`
}

func (o ListRecipientsStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecipientsStatusRequest struct{}"
	}

	return strings.Join([]string{"ListRecipientsStatusRequest", string(data)}, " ")
}
