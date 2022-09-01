package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StacksConfig struct {
	Attributes *StacksAttribute `json:"attributes,omitempty" xml:"attributes"`

	Recipe *Recipe `json:"recipe,omitempty" xml:"recipe"`
}

func (o StacksConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StacksConfig struct{}"
	}

	return strings.Join([]string{"StacksConfig", string(data)}, " ")
}
