package model

import (
	"encoding/json"

	"strings"
)

//
type LimitSpeedReq struct {
	// 任务id

	JobId string `json:"job_id"`
	// 限速信息请求体

	SpeedLimit []SpeedLimitInfo `json:"speed_limit"`
}

func (o LimitSpeedReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LimitSpeedReq struct{}"
	}

	return strings.Join([]string{"LimitSpeedReq", string(data)}, " ")
}
