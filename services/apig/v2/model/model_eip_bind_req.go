package model

import (
	"encoding/json"

	"strings"
)

type EipBindReq struct {
	// 弹性公网IP编号

	EipId *string `json:"eip_id,omitempty"`
}

func (o EipBindReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EipBindReq struct{}"
	}

	return strings.Join([]string{"EipBindReq", string(data)}, " ")
}
