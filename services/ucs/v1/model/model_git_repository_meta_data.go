package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GitRepositoryMetaData struct {

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 所属命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 唯一标识符
	Uid *string `json:"uid,omitempty"`

	// 资源的内部版本标识，用于并发控制
	ResourceVersion *string `json:"resourceVersion,omitempty"`

	// 资源的期望状态的代数，每次修改spec会自增
	Generation *int32 `json:"generation,omitempty"`

	// 创建时间
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// 删除前需要执行的清理操作
	Finalizers *[]string `json:"finalizers,omitempty"`

	// 用于跟踪资源字段管理权，记录每个字段的管理者
	ManagedFields *[]ManagedFieldsEntry `json:"managedFields,omitempty"`
}

func (o GitRepositoryMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GitRepositoryMetaData struct{}"
	}

	return strings.Join([]string{"GitRepositoryMetaData", string(data)}, " ")
}
