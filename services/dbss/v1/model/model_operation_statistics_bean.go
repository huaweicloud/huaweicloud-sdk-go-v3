package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationStatisticsBean struct {

	// 周期
	Period *string `json:"period,omitempty"`

	// 风险操作数量
	RiskOperationNum *int64 `json:"risk_operation_num,omitempty"`
}

func (o OperationStatisticsBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationStatisticsBean struct{}"
	}

	return strings.Join([]string{"OperationStatisticsBean", string(data)}, " ")
}
