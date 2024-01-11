package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectIdInfo struct {

	// 作业算子名称
	Name *string `json:"name,omitempty"`

	// 资产类型
	TypeName *string `json:"type_name,omitempty"`

	// 作业资产唯一限定名称
	QualifiedName *string `json:"qualified_name,omitempty"`
}

func (o ObjectIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectIdInfo struct{}"
	}

	return strings.Join([]string{"ObjectIdInfo", string(data)}, " ")
}
