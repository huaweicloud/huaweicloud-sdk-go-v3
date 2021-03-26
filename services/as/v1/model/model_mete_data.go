package model

import (
	"encoding/json"

	"strings"
)

// 用户自定义键值对
type MeteData struct {
	// metadata键、值。

	UserDefinedKeyValuePairs *string `json:"User-defined_key-value_pairs,omitempty"`
}

func (o MeteData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MeteData struct{}"
	}

	return strings.Join([]string{"MeteData", string(data)}, " ")
}
