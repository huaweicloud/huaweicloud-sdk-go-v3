package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// used_rds参数说明
type DatabaseInstabcesParam struct {

	// 逻辑库关联的RDS的id
	Id string `json:"id" xml:"id"`

	// 关联RDS实例的用户。
	AdminUser string `json:"adminUser" xml:"adminUser"`

	// 关联RDS实例的密码。
	AdminPassword string `json:"adminPassword" xml:"adminPassword"`
}

func (o DatabaseInstabcesParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInstabcesParam struct{}"
	}

	return strings.Join([]string{"DatabaseInstabcesParam", string(data)}, " ")
}
