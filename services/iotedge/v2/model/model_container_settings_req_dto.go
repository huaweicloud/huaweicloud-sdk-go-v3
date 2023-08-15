package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerSettingsReqDto struct {
	Configs *ContainerConfigsReqDto `json:"configs,omitempty"`
}

func (o ContainerSettingsReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerSettingsReqDto struct{}"
	}

	return strings.Join([]string{"ContainerSettingsReqDto", string(data)}, " ")
}
