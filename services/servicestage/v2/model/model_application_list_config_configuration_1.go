package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationListConfigConfiguration1 struct {
	// 应用ID。

	ApplicationId *string `json:"application_id,omitempty"`
	// 环境ID。

	EnvironmentId *string `json:"environment_id,omitempty"`

	Configuration *ApplicationListConfigConfiguration `json:"configuration,omitempty"`
}

func (o ApplicationListConfigConfiguration1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationListConfigConfiguration1 struct{}"
	}

	return strings.Join([]string{"ApplicationListConfigConfiguration1", string(data)}, " ")
}
