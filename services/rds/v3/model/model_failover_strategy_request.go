package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailoverStrategyRequest struct {
	// 可用性策略

	RepairStrategy string `json:"repairStrategy"`
}

func (o FailoverStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverStrategyRequest struct{}"
	}

	return strings.Join([]string{"FailoverStrategyRequest", string(data)}, " ")
}
