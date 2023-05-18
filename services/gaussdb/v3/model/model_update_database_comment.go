package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库备注信息。
type UpdateDatabaseComment struct {

	// 数据库名称。
	Name string `json:"name"`

	// 数据库备注,长度不能超过512个字符，不能包含回车和特殊字符!<\"='>&。
	Comment string `json:"comment"`
}

func (o UpdateDatabaseComment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseComment struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseComment", string(data)}, " ")
}
