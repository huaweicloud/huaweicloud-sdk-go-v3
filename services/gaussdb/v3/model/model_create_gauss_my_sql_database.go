package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库信息列表，列表最大长度为50。
type CreateGaussMySqlDatabase struct {

	// 数据库名称,数据库名称长度可在1～64个字符之间，由字母、数字、下划线、中划线组成，中划线的累计总长度小于等于10个字符，且不能包含其他特殊字符。
	Name string `json:"name"`

	// 数据库备注,长度不能超过512个字符，不能包含回车和特殊字符!<\"='>&。
	Comment *string `json:"comment,omitempty"`

	// 数据库使用的字符集名称，如utf8mb4、gbk。
	CharacterSet string `json:"character_set"`

	// 数据库用户列表，即创建数据库时同步授权给列表中的用户，列表最大长度为50。列表可以为空，即创建数据库时不授予其权限到数据库用户，在需要给该数据库授权某数据库用户时，调用数据库用户授权接口即可。
	Users *[]GaussMySqlDatabaseUser `json:"users,omitempty"`
}

func (o CreateGaussMySqlDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlDatabase struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlDatabase", string(data)}, " ")
}
