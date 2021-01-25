package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProtectableResponse struct {
	// 可保护性查询实例
	Instances      *[]ProtectablesResp `json:"instances,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListProtectableResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProtectableResponse struct{}"
	}

	return strings.Join([]string{"ListProtectableResponse", string(data)}, " ")
}
