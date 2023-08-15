package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceGroupResponse Response Object
type ShowDeviceGroupResponse struct {

	// 设备组ID，用于唯一标识一个设备组，在创建设备组时由物联网平台分配。
	GroupId *string `json:"group_id,omitempty"`

	// 设备组名称，单个资源空间下不可重复。
	Name *string `json:"name,omitempty"`

	// 设备组描述。
	Description *string `json:"description,omitempty"`

	// 父设备组ID，该设备组的父设备组ID。
	SuperGroupId *string `json:"super_group_id,omitempty"`

	// 设备组类型，分为动态设备组和静态设备组两种
	GroupType *string `json:"group_type,omitempty"`

	// 动态设备组规则
	DynamicGroupRule *string `json:"dynamic_group_rule,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowDeviceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupResponse", string(data)}, " ")
}
