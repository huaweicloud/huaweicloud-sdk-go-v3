package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationListConfigConfigurationEnv struct {

	// 环境变量名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 环境变量取值。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o ApplicationListConfigConfigurationEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationListConfigConfigurationEnv struct{}"
	}

	return strings.Join([]string{"ApplicationListConfigConfigurationEnv", string(data)}, " ")
}
