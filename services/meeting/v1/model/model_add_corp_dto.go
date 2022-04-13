package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddCorpDto struct {
	BasicInfo *CorpBasicDto `json:"basicInfo"`

	AdminInfo *AdminDto `json:"adminInfo"`

	ResInfo *AddCorpResDto `json:"resInfo,omitempty"`
	// 媒体接入（包括SBC和MCU）分组id, 可通过企业资源管理下的SP管理员查询资源信息接口获取。

	GroupId *string `json:"groupId,omitempty"`
	// 可配置项信息。

	PropertyInfo *[]OrgPropertyDto `json:"propertyInfo,omitempty"`
}

func (o AddCorpDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCorpDto struct{}"
	}

	return strings.Join([]string{"AddCorpDto", string(data)}, " ")
}
