package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateInstanceReq struct {
	Instance *CreateInstanceDetail `json:"instance"`

	ExtendParam *CreateInstanceExtendParam `json:"extend_param,omitempty"`
}

func (o CreateInstanceReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceReq", string(data)}, " ")
}
