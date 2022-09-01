package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建容灾演练请求体
type CreateDisasterRecoveryDrillRequestBody struct {
	DisasterRecoveryDrill *CreateDisasterRecoveryDrillRequestParams `json:"disaster_recovery_drill" xml:"disaster_recovery_drill"`
}

func (o CreateDisasterRecoveryDrillRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryDrillRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryDrillRequestBody", string(data)}, " ")
}
