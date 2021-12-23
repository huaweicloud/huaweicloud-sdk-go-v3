package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Ignore struct {
	// API请求方法

	Method *string `json:"method,omitempty"`
	// API请求路径

	Path *string `json:"path,omitempty"`
}

func (o Ignore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ignore struct{}"
	}

	return strings.Join([]string{"Ignore", string(data)}, " ")
}
