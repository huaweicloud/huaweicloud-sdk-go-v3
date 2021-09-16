package model

import (
	"encoding/json"

	"strings"
)

// 切换边缘实例的元数据
type ChangeOsMetadata struct {
	// 切换边缘实例操作系统过程中注入的用户数据。

	UserData *string `json:"user_data,omitempty"`
}

func (o ChangeOsMetadata) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeOsMetadata struct{}"
	}

	return strings.Join([]string{"ChangeOsMetadata", string(data)}, " ")
}
