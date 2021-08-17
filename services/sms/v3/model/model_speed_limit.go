package model

import (
	"encoding/json"

	"strings"
)

// 修改速率的参数
type SpeedLimit struct {
	// 按时间段限速信息

	SpeedLimit []SpeedLimitlJson `json:"speed_limit"`
}

func (o SpeedLimit) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SpeedLimit struct{}"
	}

	return strings.Join([]string{"SpeedLimit", string(data)}, " ")
}
