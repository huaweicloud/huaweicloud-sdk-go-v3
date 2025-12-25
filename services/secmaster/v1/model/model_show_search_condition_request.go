package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchConditionRequest Request Object
type ShowSearchConditionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 条件ID
	ConditionId string `json:"condition_id"`
}

func (o ShowSearchConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchConditionRequest struct{}"
	}

	return strings.Join([]string{"ShowSearchConditionRequest", string(data)}, " ")
}
