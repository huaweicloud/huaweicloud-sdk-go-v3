package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSearchConditionRequest Request Object
type UpdateSearchConditionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 条件ID
	ConditionId string `json:"condition_id"`

	Body *UpdateSearchConditionRequestBody `json:"body,omitempty"`
}

func (o UpdateSearchConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSearchConditionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSearchConditionRequest", string(data)}, " ")
}
