package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationConfigModify struct {
	// 环境ID。

	EnvironmentId string `json:"environment_id"`

	Configuration *ApplicationConfigModifyConfiguration `json:"configuration"`
}

func (o ApplicationConfigModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationConfigModify struct{}"
	}

	return strings.Join([]string{"ApplicationConfigModify", string(data)}, " ")
}
