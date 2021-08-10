package model

import (
	"encoding/json"

	"strings"
)

//
type SubnetList struct {
	// 路由表关联的子网ID

	Id string `json:"id"`
}

func (o SubnetList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SubnetList struct{}"
	}

	return strings.Join([]string{"SubnetList", string(data)}, " ")
}
