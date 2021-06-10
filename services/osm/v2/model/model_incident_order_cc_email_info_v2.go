package model

import (
	"encoding/json"

	"strings"
)

type IncidentOrderCcEmailInfoV2 struct {
	// 用户id

	UserId *string `json:"user_id,omitempty"`
	// 客户id

	CustomerId *string `json:"customer_id,omitempty"`
	// 操作邮箱

	CcEmail *string `json:"cc_email,omitempty"`
}

func (o IncidentOrderCcEmailInfoV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IncidentOrderCcEmailInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentOrderCcEmailInfoV2", string(data)}, " ")
}
