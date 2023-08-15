package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Grant 授权信息
type Grant struct {

	// 权限名称，根据不通数据库对象类型，拥有权限不同 - database  CREATE | CONNECT | TEMPORARY | TEMP  ALL  PRIVILEGES - schema  CREATE | USAGE | ALTER | DROP  ALL  PRIVILEGES - table  SELECT | INSERT | UPDATE | DELETE | TRUNCATE | REFERENCES | TRIGGER | ANALYZE | ANALYSE | VACUUM | ALTER | DROP  ALL  PRIVILEGES - view  SELECT | INSERT | UPDATE | DELETE | TRUNCATE | REFERENCES   | TRIGGER | ANALYZE | ANALYSE | VACUUM | ALTER | DROP  ALL  PRIVILEGES - column  SELECT | INSERT | UPDATE | REFERENCES  ALL  PRIVILEGES - function  EXECUTE  ALL  PRIVILEGES - sequence  SELECT | UPDATE | USAGE  ALL  PRIVILEGES - nodegroup  CREATE | USAGE | COMPUTE  ALL  PRIVILEGES - role  role_name（角色名称）
	Permission string `json:"permission"`

	// 是否包含授权选项
	GrantWith bool `json:"grant_with"`
}

func (o Grant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Grant struct{}"
	}

	return strings.Join([]string{"Grant", string(data)}, " ")
}
