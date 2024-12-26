package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowDbRolePrivilegesRequestBody 授权数据库角色。
type AllowDbRolePrivilegesRequestBody struct {

	// 数据库名称。 不能使用模板库，且是已存在的数据库名称。 模板库包括postgres， template0 ，template1，templatea，template_pdb，templatem。
	DbName string `json:"db_name"`

	User *GaussDBforOpenGaussRoleAttributes `json:"user"`
}

func (o AllowDbRolePrivilegesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbRolePrivilegesRequestBody struct{}"
	}

	return strings.Join([]string{"AllowDbRolePrivilegesRequestBody", string(data)}, " ")
}
