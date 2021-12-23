package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationModify struct {
	// 应用名称。

	Name *string `json:"name,omitempty"`
	// 应用描述。

	Description *string `json:"description,omitempty"`
}

func (o ApplicationModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationModify struct{}"
	}

	return strings.Join([]string{"ApplicationModify", string(data)}, " ")
}
