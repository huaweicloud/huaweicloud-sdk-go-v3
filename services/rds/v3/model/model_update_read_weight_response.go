package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateReadWeightResponse struct {
	// 修改读写分离权重或延时阈值的结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReadWeightResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateReadWeightResponse struct{}"
	}

	return strings.Join([]string{"UpdateReadWeightResponse", string(data)}, " ")
}
