package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchConditionRequest Request Object
type CreateSearchConditionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateSearchConditionRequestBody `json:"body,omitempty"`
}

func (o CreateSearchConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchConditionRequest struct{}"
	}

	return strings.Join([]string{"CreateSearchConditionRequest", string(data)}, " ")
}
