package model

import (
	"encoding/json"

	"strings"
)

type SysTags struct {
	// 企业项目ID

	Value string `json:"value"`
	// 该值目前固定为“_sys_enterprise_project_id”

	Key string `json:"key"`
}

func (o SysTags) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SysTags struct{}"
	}

	return strings.Join([]string{"SysTags", string(data)}, " ")
}
