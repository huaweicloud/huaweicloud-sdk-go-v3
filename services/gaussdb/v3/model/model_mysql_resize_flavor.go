package model

import (
	"encoding/json"

	"strings"
)

type MysqlResizeFlavor struct {
	// 规格码

	SpecCode string `json:"spec_code"`
}

func (o MysqlResizeFlavor) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlResizeFlavor struct{}"
	}

	return strings.Join([]string{"MysqlResizeFlavor", string(data)}, " ")
}
