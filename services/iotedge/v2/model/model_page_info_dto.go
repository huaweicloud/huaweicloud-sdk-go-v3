package model

import (
	"encoding/json"

	"strings"
)

type PageInfoDto struct {
	// 页码

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数

	Limit *int32 `json:"limit,omitempty"`
}

func (o PageInfoDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PageInfoDto struct{}"
	}

	return strings.Join([]string{"PageInfoDto", string(data)}, " ")
}
