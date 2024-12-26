package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenGaussRoleAttributes 用户的权限属性。
type GaussDBforOpenGaussRoleAttributes struct {

	// 数据库角色名称。 不能使用系统用户或角色，且名称必须存在。 系统用户/角色包括“rdsAdmin”,“rdsMetric”, “rdsBackup”, “rdsRepl”, “root”。
	Name string `json:"name"`

	// SCHEMA名称。 不能和模板库以及系统内schema重名，且schema名称必须存在。 模板库包括postgres， template0 ，template1, 系统内schema包括public，information_schema。
	Schema string `json:"schema"`

	// 数据库角色权限。 - true：只读。 - false：可读可写。
	Readonly bool `json:"readonly"`

	// 数据库用户/角色名称。 该字段的含义是将此用户/角色的权限授予给name字段指定的角色，通过readonly字段判断是否授予只读权限。 不能和系统用户/角色名称相同，且用户/角色名称必须存在，系统用户/角色包括“rdsAdmin”,“ rdsMetric”, “rdsBackup”, “rdsRepl”, “root”。
	DefaultPrivilegeGrantee *string `json:"default_privilege_grantee,omitempty"`
}

func (o GaussDBforOpenGaussRoleAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenGaussRoleAttributes struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenGaussRoleAttributes", string(data)}, " ")
}
