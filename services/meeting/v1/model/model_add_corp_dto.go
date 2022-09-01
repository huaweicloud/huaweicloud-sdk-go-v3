package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddCorpDto struct {
	BasicInfo *CorpBasicDto `json:"basicInfo" xml:"basicInfo"`

	AdminInfo *AdminDto `json:"adminInfo" xml:"adminInfo"`

	ResInfo *AddCorpResDto `json:"resInfo,omitempty" xml:"resInfo"`

	// 媒体接入（包括SBC和MCU）分组id, 可通过企业资源管理下的SP管理员查询资源信息接口获取。
	GroupId *string `json:"groupId,omitempty" xml:"groupId"`

	// 可配置项信息。
	PropertyInfo *[]OrgPropertyDto `json:"propertyInfo,omitempty" xml:"propertyInfo"`
}

func (o AddCorpDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCorpDto struct{}"
	}

	return strings.Join([]string{"AddCorpDto", string(data)}, " ")
}
