package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataConsumptionRequest Request Object
type ShowDataConsumptionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`
}

func (o ShowDataConsumptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataConsumptionRequest struct{}"
	}

	return strings.Join([]string{"ShowDataConsumptionRequest", string(data)}, " ")
}
