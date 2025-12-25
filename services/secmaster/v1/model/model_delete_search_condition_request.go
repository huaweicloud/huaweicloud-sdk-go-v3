package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSearchConditionRequest Request Object
type DeleteSearchConditionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 条件ID
	ConditionId string `json:"condition_id"`
}

func (o DeleteSearchConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSearchConditionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSearchConditionRequest", string(data)}, " ")
}
