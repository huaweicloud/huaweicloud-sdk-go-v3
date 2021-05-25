package model

import (
	"encoding/json"

	"strings"
)

// 函数异步调用目标参数配置。
type FuncAsyncDestinationConfig struct {
	OnSuccess *FuncDestinationConfig `json:"on_success,omitempty"`

	OnFailure *FuncDestinationConfig `json:"on_failure,omitempty"`
}

func (o FuncAsyncDestinationConfig) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FuncAsyncDestinationConfig struct{}"
	}

	return strings.Join([]string{"FuncAsyncDestinationConfig", string(data)}, " ")
}
