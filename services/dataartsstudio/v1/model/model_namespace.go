package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Namespace struct {

	// 命名空间的名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 命名空间的guid
	NamespaceGuid *string `json:"namespace_guid,omitempty"`

	// 命名空间的唯一标识名称
	NamespaceQualifiedName *string `json:"namespace_qualified_name,omitempty"`

	// 命名空间下的表总数
	TableCount *int32 `json:"table_count,omitempty"`
}

func (o Namespace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Namespace struct{}"
	}

	return strings.Join([]string{"Namespace", string(data)}, " ")
}
