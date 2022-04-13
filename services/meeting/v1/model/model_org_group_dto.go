package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 媒体接入(包括SBC和MCU)分组信息
type OrgGroupDto struct {
	// 分组Id

	GroupId *string `json:"groupId,omitempty"`
	// 分组名称

	GroupName *string `json:"groupName,omitempty"`
}

func (o OrgGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrgGroupDto struct{}"
	}

	return strings.Join([]string{"OrgGroupDto", string(data)}, " ")
}
