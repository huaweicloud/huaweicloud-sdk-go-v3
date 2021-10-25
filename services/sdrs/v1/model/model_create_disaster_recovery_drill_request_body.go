package model

import (
	"encoding/json"

	"strings"
)

// 创建容灾演练请求体
type CreateDisasterRecoveryDrillRequestBody struct {
	DisasterRecoveryDrill *CreateDisasterRecoveryDrillRequestParams `json:"disaster_recovery_drill"`
}

func (o CreateDisasterRecoveryDrillRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryDrillRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryDrillRequestBody", string(data)}, " ")
}
