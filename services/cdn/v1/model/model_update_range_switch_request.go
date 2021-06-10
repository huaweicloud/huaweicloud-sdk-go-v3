package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRangeSwitchRequest struct {
	// 加速域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`

	Body *RangeStatusRequest `json:"body,omitempty"`
}

func (o UpdateRangeSwitchRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRangeSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateRangeSwitchRequest", string(data)}, " ")
}
