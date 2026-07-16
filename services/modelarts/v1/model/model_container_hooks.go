package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerHooks struct {
	PostStart *Config `json:"post_start,omitempty"`

	PreStart *Config `json:"pre_start,omitempty"`
}

func (o ContainerHooks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerHooks struct{}"
	}

	return strings.Join([]string{"ContainerHooks", string(data)}, " ")
}
