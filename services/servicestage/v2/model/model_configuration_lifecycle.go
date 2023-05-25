package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 生命周期
type ConfigurationLifecycle struct {
	Entrypoint *LifecycleEntrypoint `json:"entrypoint,omitempty"`

	PostStart *LifecycleEntrypoint `json:"post-start,omitempty"`

	PreStop *LifecycleEntrypoint `json:"pre-stop,omitempty"`
}

func (o ConfigurationLifecycle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationLifecycle struct{}"
	}

	return strings.Join([]string{"ConfigurationLifecycle", string(data)}, " ")
}
