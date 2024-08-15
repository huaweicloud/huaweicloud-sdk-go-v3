package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleGroupInfo struct {

	// 角色id
	RoleId string `json:"role_id"`

	// 场景id
	SceneId string `json:"scene_id"`
}

func (o ScheduleGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleGroupInfo struct{}"
	}

	return strings.Join([]string{"ScheduleGroupInfo", string(data)}, " ")
}
