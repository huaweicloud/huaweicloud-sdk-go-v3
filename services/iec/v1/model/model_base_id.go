package model

import (
	"encoding/json"

	"strings"
)

// ID对象
type BaseId struct {
	// 对象ID，uuid。

	Id string `json:"id"`
}

func (o BaseId) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BaseId struct{}"
	}

	return strings.Join([]string{"BaseId", string(data)}, " ")
}
