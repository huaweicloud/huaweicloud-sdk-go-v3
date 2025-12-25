package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDataConsumptionRequest Request Object
type EnableDataConsumptionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`
}

func (o EnableDataConsumptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataConsumptionRequest struct{}"
	}

	return strings.Join([]string{"EnableDataConsumptionRequest", string(data)}, " ")
}
