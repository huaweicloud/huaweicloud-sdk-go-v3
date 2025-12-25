package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableDataConsumptionRequest Request Object
type DisableDataConsumptionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 表ID
	TableId string `json:"table_id"`

	Body *DisableConsumptionV2RequestBody `json:"body,omitempty"`
}

func (o DisableDataConsumptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDataConsumptionRequest struct{}"
	}

	return strings.Join([]string{"DisableDataConsumptionRequest", string(data)}, " ")
}
