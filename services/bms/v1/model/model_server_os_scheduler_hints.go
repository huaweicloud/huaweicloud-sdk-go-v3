package model

import (
	"encoding/json"

	"strings"
)

type ServerOsSchedulerHints struct {
	// 反亲和性组信息。  UUID格式。

	Group *[]string `json:"group,omitempty"`
}

func (o ServerOsSchedulerHints) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServerOsSchedulerHints struct{}"
	}

	return strings.Join([]string{"ServerOsSchedulerHints", string(data)}, " ")
}
