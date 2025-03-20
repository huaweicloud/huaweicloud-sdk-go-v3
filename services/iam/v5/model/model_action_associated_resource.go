package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionAssociatedResource 与该授权项关联的资源。
type ActionAssociatedResource struct {

	// 统一资源名称模板，表示可以通过这类资源的统一资源名称对该授权项进行资源粒度的授权。
	UrnTemplate string `json:"urn_template"`

	// 标识该资源类型是否是这个授权项必选的，即授权项一定涉及对这类资源的操作；例如subnet是vpc:subnets:get的必选资源类型；而ou是organizations::tagResource的可选资源类型，因为organizations::tagResource操作的资源还可能是account或者policy。
	Required bool `json:"required"`

	// 针对该授权项和资源的服务自定义条件属性以及部分全局属性，只有授权项和资源同时匹配时才会生效。
	ConditionKeys *[]string `json:"condition_keys,omitempty"`
}

func (o ActionAssociatedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionAssociatedResource struct{}"
	}

	return strings.Join([]string{"ActionAssociatedResource", string(data)}, " ")
}
