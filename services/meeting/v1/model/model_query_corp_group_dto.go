package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCorpGroupDto struct {

	// 分组Id
	GroupId *string `json:"groupId,omitempty" xml:"groupId"`

	// 分组名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName"`

	// 分组类型
	GroupType *int32 `json:"groupType,omitempty" xml:"groupType"`

	// 分组备注信息
	Remarks *string `json:"remarks,omitempty" xml:"remarks"`

	// 区域ID，仅服务列表类型场景必填
	RegionId *string `json:"regionId,omitempty" xml:"regionId"`

	// 分组状态 - 0: 正常 - 1: 停用，服务列表类型停用后创建企业就不会再自动分配到该分组
	Status *int32 `json:"status,omitempty" xml:"status"`
}

func (o QueryCorpGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpGroupDto struct{}"
	}

	return strings.Join([]string{"QueryCorpGroupDto", string(data)}, " ")
}
