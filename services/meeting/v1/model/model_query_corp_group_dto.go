package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCorpGroupDto struct {

	// 媒体接入分组id。
	GroupId *string `json:"groupId,omitempty"`

	// 分组名称。
	GroupName *string `json:"groupName,omitempty"`

	// 分组类型。
	GroupType *int32 `json:"groupType,omitempty"`

	// 分组备注信息。
	Remarks *string `json:"remarks,omitempty"`

	// 区域ID，仅服务列表类型场景必填。
	RegionId *string `json:"regionId,omitempty"`

	// 分组状态。 - 0: 正常 - 1: 停用，服务列表类型停用后创建企业就不会再自动分配到该分组
	Status *int32 `json:"status,omitempty"`
}

func (o QueryCorpGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpGroupDto struct{}"
	}

	return strings.Join([]string{"QueryCorpGroupDto", string(data)}, " ")
}
