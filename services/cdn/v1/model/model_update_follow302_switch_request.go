package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateFollow302SwitchRequest struct {
	// 加速域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`

	Body *Follow302StatusRequest `json:"body,omitempty"`
}

func (o UpdateFollow302SwitchRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFollow302SwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateFollow302SwitchRequest", string(data)}, " ")
}
