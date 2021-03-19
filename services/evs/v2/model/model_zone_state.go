package model

import (
	"encoding/json"

	"strings"
)

// 可用分区的状态。
type ZoneState struct {
	// 可用分区是否可用。

	Available *bool `json:"available,omitempty"`
}

func (o ZoneState) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ZoneState struct{}"
	}

	return strings.Join([]string{"ZoneState", string(data)}, " ")
}
