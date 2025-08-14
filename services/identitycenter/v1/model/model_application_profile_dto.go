package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationProfileDto 应用程序Profile
type ApplicationProfileDto struct {

	// 关联关系名称，默认为Default
	Name string `json:"name"`

	// 关联关系状态，默认为ENABLED
	Status string `json:"status"`

	// 实体与应用程序关联ID
	ProfileId string `json:"profile_id"`
}

func (o ApplicationProfileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationProfileDto struct{}"
	}

	return strings.Join([]string{"ApplicationProfileDto", string(data)}, " ")
}
