package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModCorpDto struct {
	BasicInfo *ModCorpBasicDto `json:"basicInfo,omitempty" xml:"basicInfo"`

	AdminInfo *ModAdminDto `json:"adminInfo,omitempty" xml:"adminInfo"`

	// 媒体接入（包括SBC和MCU）分组id, 可通过企业资源管理下的SP管理员查询资源信息接口获取。
	GroupId *string `json:"groupId,omitempty" xml:"groupId"`

	// 可配置项信息。
	PropertyInfo *[]OrgPropertyDto `json:"propertyInfo,omitempty" xml:"propertyInfo"`
}

func (o ModCorpDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModCorpDto struct{}"
	}

	return strings.Join([]string{"ModCorpDto", string(data)}, " ")
}
