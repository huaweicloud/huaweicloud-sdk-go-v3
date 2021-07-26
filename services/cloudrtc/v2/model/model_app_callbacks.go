package model

import (
	"encoding/json"

	"strings"
)

// app回调配置
type AppCallbacks struct {
	PushCallback *AppCallbackUrl `json:"push_callback,omitempty"`

	RecordCallback *AppCallbackUrl `json:"record_callback,omitempty"`
}

func (o AppCallbacks) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AppCallbacks struct{}"
	}

	return strings.Join([]string{"AppCallbacks", string(data)}, " ")
}
