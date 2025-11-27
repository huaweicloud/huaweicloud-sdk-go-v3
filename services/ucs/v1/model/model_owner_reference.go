package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OwnerReference 用于在Kubernetes中标识一个拥有者对象，它提供了足够的信息来识别和关联资源之间的所有权关系
type OwnerReference struct {

	// 标识引用对象的API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	// 引用对象的类型
	Kind *string `json:"kind,omitempty"`

	// 引用对象的名称
	Name *string `json:"name,omitempty"`

	// 引用对象的uid
	Uid *string `json:"uid,omitempty"`

	// 如果为true，表示该引用指向管理该资源的控制器
	Controller *bool `json:"controller,omitempty"`

	// 当为true且拥有者有名为\"foregroundDeletion\"的finalizer 时，会阻止拥有者被删除，直到这个引用被移除
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty"`
}

func (o OwnerReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OwnerReference struct{}"
	}

	return strings.Join([]string{"OwnerReference", string(data)}, " ")
}
