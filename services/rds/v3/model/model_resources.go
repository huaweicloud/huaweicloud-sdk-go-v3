package model

import (
	"encoding/json"

	"strings"
)

type Resources struct {
	// 项目资源配额。

	Quota *int32 `json:"quota,omitempty"`
	// 已使用的资源数量。

	Used *int32 `json:"used,omitempty"`
	// 项目资源类型，取值范围：instance。

	Type *string `json:"type,omitempty"`
}

func (o Resources) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
