package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModuleContainerSettingsResDto struct {
	Configs *ContainerConfigsResDto `json:"configs,omitempty"`
}

func (o ModuleContainerSettingsResDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleContainerSettingsResDto struct{}"
	}

	return strings.Join([]string{"ModuleContainerSettingsResDto", string(data)}, " ")
}
