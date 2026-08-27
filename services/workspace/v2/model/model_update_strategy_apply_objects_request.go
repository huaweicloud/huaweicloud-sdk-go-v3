package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStrategyApplyObjectsRequest Request Object
type UpdateStrategyApplyObjectsRequest struct {

	// 策略ID
	StrategyId string `json:"strategy_id"`

	Body *UpdateStrategyApplyObjectsRequestBody `json:"body,omitempty"`
}

func (o UpdateStrategyApplyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStrategyApplyObjectsRequest struct{}"
	}

	return strings.Join([]string{"UpdateStrategyApplyObjectsRequest", string(data)}, " ")
}
