package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StacksConfig struct {
	Attributes *StacksAttribute `json:"attributes,omitempty"`

	Recipe *Recipe `json:"recipe,omitempty"`
}

func (o StacksConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StacksConfig struct{}"
	}

	return strings.Join([]string{"StacksConfig", string(data)}, " ")
}
