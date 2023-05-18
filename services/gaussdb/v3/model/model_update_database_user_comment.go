package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库用户备注信息。
type UpdateDatabaseUserComment struct {

	// 数据库用户名。
	Name string `json:"name"`

	// 主机地址。
	Host string `json:"host"`

	// 数据库用户备注,长度不能超过512个字符，不能包含回车和特殊字符!<\"='>&。  该字段只针对新版本的实例生效，必须大于等于指定的内核版本-2.0.13.0，如果不符合内核版本要求，参考升级内核版本升级至最新。
	Comment string `json:"comment"`
}

func (o UpdateDatabaseUserComment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseUserComment struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseUserComment", string(data)}, " ")
}
