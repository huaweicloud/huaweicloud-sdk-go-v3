package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceGroup struct {

	// 实例组ID
	GroupId int32 `json:"group_id"`

	// 实例组名称
	GroupName string `json:"group_name"`

	// 描述
	Description string `json:"description"`
}

func (o InstanceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceGroup struct{}"
	}

	return strings.Join([]string{"InstanceGroup", string(data)}, " ")
}
