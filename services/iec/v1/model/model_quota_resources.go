package model

import (
	"encoding/json"

	"strings"
)

//
type QuotaResources struct {
	// 配额信息列表。

	Resources *[]QuotaResource `json:"resources,omitempty"`
}

func (o QuotaResources) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaResources struct{}"
	}

	return strings.Join([]string{"QuotaResources", string(data)}, " ")
}
