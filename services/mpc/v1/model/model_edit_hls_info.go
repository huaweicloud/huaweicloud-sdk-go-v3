package model

import (
	"encoding/json"

	"strings"
)

type EditHlsInfo struct {
	// 切片间隔。

	Interval *int32 `json:"interval,omitempty"`
}

func (o EditHlsInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EditHlsInfo struct{}"
	}

	return strings.Join([]string{"EditHlsInfo", string(data)}, " ")
}
