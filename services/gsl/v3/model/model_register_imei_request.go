package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RegisterImeiRequest struct {
	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取

	SimCardId int64 `json:"sim_card_id"`

	Body *RegisterImeiReq `json:"body,omitempty"`
}

func (o RegisterImeiRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterImeiRequest struct{}"
	}

	return strings.Join([]string{"RegisterImeiRequest", string(data)}, " ")
}
