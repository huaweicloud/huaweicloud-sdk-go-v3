package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModCorpDto struct {
	BasicInfo *ModCorpBasicDto `json:"basicInfo,omitempty"`

	AdminInfo *ModAdminDto `json:"adminInfo,omitempty"`

	// 媒体接入（包括SBC和MCU）分组id, 可通过企业资源管理下的SP管理员查询资源信息接口获取。
	GroupId *string `json:"groupId,omitempty"`

	// 可配置项信息。
	PropertyInfo *[]OrgPropertyDto `json:"propertyInfo,omitempty"`
}

func (o ModCorpDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModCorpDto struct{}"
	}

	return strings.Join([]string{"ModCorpDto", string(data)}, " ")
}
