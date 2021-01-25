package model

import (
	"encoding/json"

	"strings"
)

type PolicyCreateReq struct {
	Policy *PolicyCreate `json:"policy"`
}

func (o PolicyCreateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PolicyCreateReq struct{}"
	}

	return strings.Join([]string{"PolicyCreateReq", string(data)}, " ")
}
