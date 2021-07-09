package model

import (
	"encoding/json"

	"strings"
)

type ProjectQuotas struct {
	// 资源列表对象。

	Resources []Resource `json:"resources"`
}

func (o ProjectQuotas) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProjectQuotas struct{}"
	}

	return strings.Join([]string{"ProjectQuotas", string(data)}, " ")
}
