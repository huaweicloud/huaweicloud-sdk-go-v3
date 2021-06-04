package model

import (
	"encoding/json"

	"strings"
)

// 修改设备影子预期数据结构体。
type UpdateDesireds struct {
	// **参数说明**：设备影子期望值构体。

	Shadow *[]UpdateDesired `json:"shadow,omitempty"`
}

func (o UpdateDesireds) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDesireds struct{}"
	}

	return strings.Join([]string{"UpdateDesireds", string(data)}, " ")
}
