package model

import (
	"encoding/json"

	"strings"
)

//
type QuotaResult struct {
	// 资源信息

	Resources *[]Resources `json:"resources,omitempty"`
}

func (o QuotaResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaResult struct{}"
	}

	return strings.Join([]string{"QuotaResult", string(data)}, " ")
}
