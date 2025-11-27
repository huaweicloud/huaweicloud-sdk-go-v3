package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectMeta ObjectMeta是Kubernetes中所有持久化资源必须包含的元数据结构
type ObjectMeta struct {

	// 资源ID
	Uid *string `json:"uid,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 当未提供name时，服务器使用的前缀来生成唯一名称
	GenerateName *string `json:"generateName,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 标签
	Labels map[string]string `json:"labels,omitempty"`

	// 注解
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// 更新时间
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	// 资源内部版本
	ResourceVersion *string `json:"resourceVersion,omitempty"`

	// 资源期望状态的代数
	Generation *string `json:"generation,omitempty"`

	// 记录哪些字段由哪些工作流管理
	ManagedFields *[]ManagedFieldsEntry `json:"managedFields,omitempty"`

	// 用于定义对象间所有权关系，管理对象的依赖关系和垃圾回收机制，支持控制器对资源的管理
	OwnerReferences *[]OwnerReference `json:"ownerReferences,omitempty"`
}

func (o ObjectMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectMeta struct{}"
	}

	return strings.Join([]string{"ObjectMeta", string(data)}, " ")
}
