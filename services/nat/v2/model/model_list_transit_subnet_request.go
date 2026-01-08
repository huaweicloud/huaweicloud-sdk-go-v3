package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitSubnetRequest Request Object
type ListTransitSubnetRequest struct {

	// 中转子网的ID
	Id *[]string `json:"id,omitempty"`

	// 中转子网的名字
	Name *[]string `json:"name,omitempty"`

	// 中转子网的描述
	Description *[]string `json:"description,omitempty"`

	// 中转子网的子网所属项目的ID
	VirsubnetProjectId *[]string `json:"virsubnet_project_id,omitempty"`

	// 中转子网的子网所属的VPC的ID
	VpcId *[]string `json:"vpc_id,omitempty"`

	// 中转子网的子网ID
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	// 中转子网状态。 取值范围： - ACTIVE： 当前资源状态正常。 - INACTIVE： 不可用。
	Status *[]string `json:"status,omitempty"`

	// 功能说明：每页返回的个数。 取值范围：1~2000。 默认值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 功能说明：分页查询起始的资源ID，为空时查询第一页。 值从上一次查询的PageInfo中的next_marker或者previous_marker中获取
	Marker *string `json:"marker,omitempty"`

	// 是否查询前一页
	PageReverse *bool `json:"page_reverse,omitempty"`
}

func (o ListTransitSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitSubnetRequest struct{}"
	}

	return strings.Join([]string{"ListTransitSubnetRequest", string(data)}, " ")
}
