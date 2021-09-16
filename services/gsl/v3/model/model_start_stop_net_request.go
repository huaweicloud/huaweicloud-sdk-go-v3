package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StartStopNetRequest struct {
	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取

	SimCardId int64 `json:"sim_card_id"`

	Body *CutNetReq `json:"body,omitempty"`
}

func (o StartStopNetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartStopNetRequest struct{}"
	}

	return strings.Join([]string{"StartStopNetRequest", string(data)}, " ")
}
