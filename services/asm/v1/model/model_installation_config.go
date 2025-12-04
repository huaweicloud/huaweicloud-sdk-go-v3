package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstallationConfig struct {
	Nodes *Selector `json:"nodes"`
}

func (o InstallationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallationConfig struct{}"
	}

	return strings.Join([]string{"InstallationConfig", string(data)}, " ")
}
