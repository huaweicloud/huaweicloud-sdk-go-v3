package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationItemDto struct {

	// 要执行的修改操作类型。 add：添加属性，replace：替换属性，remove：删除属性
	Op string `json:"op"`

	// 要修改的属性路径
	Path *string `json:"path,omitempty"`

	// 要修改的属性值。执行删除操作时不填写
	Value *interface{} `json:"value,omitempty"`
}

func (o OperationItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationItemDto struct{}"
	}

	return strings.Join([]string{"OperationItemDto", string(data)}, " ")
}
