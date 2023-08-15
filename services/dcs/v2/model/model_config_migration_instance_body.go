package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigMigrationInstanceBody 迁移任务实例信息
type ConfigMigrationInstanceBody struct {

	// Redis实例ID。（Redis类型为云服务Redis时必须填写）
	Id *string `json:"id,omitempty"`

	// Redis实例地址。（Redis类型为自建Redis时必须填写）。
	Addrs *string `json:"addrs,omitempty"`

	// Redis密码，如果设置了密码，则必须填写。
	Password *string `json:"password,omitempty"`
}

func (o ConfigMigrationInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMigrationInstanceBody struct{}"
	}

	return strings.Join([]string{"ConfigMigrationInstanceBody", string(data)}, " ")
}
