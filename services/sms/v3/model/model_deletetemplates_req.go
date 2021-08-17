package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type DeletetemplatesReq struct {
	// 需要删除的模板ID

	Ids *[]string `json:"ids,omitempty"`
}

func (o DeletetemplatesReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletetemplatesReq struct{}"
	}

	return strings.Join([]string{"DeletetemplatesReq", string(data)}, " ")
}
