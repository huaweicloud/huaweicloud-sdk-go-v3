package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationListConfigConfiguration1 struct {

	// 应用ID。
	ApplicationId *string `json:"application_id,omitempty" xml:"application_id"`

	// 环境ID。
	EnvironmentId *string `json:"environment_id,omitempty" xml:"environment_id"`

	Configuration *ApplicationListConfigConfiguration `json:"configuration,omitempty" xml:"configuration"`
}

func (o ApplicationListConfigConfiguration1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationListConfigConfiguration1 struct{}"
	}

	return strings.Join([]string{"ApplicationListConfigConfiguration1", string(data)}, " ")
}
