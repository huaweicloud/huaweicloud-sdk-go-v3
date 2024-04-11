package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CodeTableFieldValueUpdateVo 码表属性值修改内容。
type CodeTableFieldValueUpdateVo struct {

	// 新增码表属性、属性值列表。
	ToAdd *[]CodeTableFieldVo `json:"to_add,omitempty"`

	// 编辑码表属性值列表。
	ToModify *[]CodeTableFieldVo `json:"to_modify,omitempty"`

	// 删除码表属性ID列表。
	ToRemove *[]CodeTableFieldVo `json:"to_remove,omitempty"`
}

func (o CodeTableFieldValueUpdateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeTableFieldValueUpdateVo struct{}"
	}

	return strings.Join([]string{"CodeTableFieldValueUpdateVo", string(data)}, " ")
}
