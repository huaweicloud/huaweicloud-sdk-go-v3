package model

import (
	"encoding/json"

	"strings"
)

// request
type PreheatingTaskRequestPreheatingTask struct {
	// 预热urls

	Urls []string `json:"urls"`
}

func (o PreheatingTaskRequestPreheatingTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PreheatingTaskRequestPreheatingTask struct{}"
	}

	return strings.Join([]string{"PreheatingTaskRequestPreheatingTask", string(data)}, " ")
}
