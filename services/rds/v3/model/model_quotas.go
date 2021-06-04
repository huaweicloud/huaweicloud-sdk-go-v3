package model

import (
	"encoding/json"

	"strings"
)

type Quotas struct {
	// 资源列表对象。

	Resources *[]Resources `json:"resources,omitempty"`
}

func (o Quotas) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Quotas struct{}"
	}

	return strings.Join([]string{"Quotas", string(data)}, " ")
}
