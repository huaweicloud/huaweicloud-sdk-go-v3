package model

import (
	"encoding/json"

	"strings"
)

//
type Versions struct {
	// 版本的资源链接信息。

	Values []Version `json:"values"`
}

func (o Versions) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Versions struct{}"
	}

	return strings.Join([]string{"Versions", string(data)}, " ")
}
