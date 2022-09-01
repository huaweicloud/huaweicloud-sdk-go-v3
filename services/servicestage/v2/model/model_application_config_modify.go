package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationConfigModify struct {

	// 环境ID。
	EnvironmentId string `json:"environment_id" xml:"environment_id"`

	Configuration *ApplicationConfigModifyConfiguration `json:"configuration" xml:"configuration"`
}

func (o ApplicationConfigModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationConfigModify struct{}"
	}

	return strings.Join([]string{"ApplicationConfigModify", string(data)}, " ")
}
