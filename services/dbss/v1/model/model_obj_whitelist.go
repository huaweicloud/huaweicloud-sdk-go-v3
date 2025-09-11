package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjWhitelist 添加白名单对象
type ObjWhitelist struct {

	// 数据库ID
	DbIds *[]string `json:"db_ids,omitempty"`

	// 描述信息
	Desc *string `json:"desc,omitempty"`

	// 状态
	Enabled *bool `json:"enabled,omitempty"`

	// SQL语句
	Sql string `json:"sql"`
}

func (o ObjWhitelist) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjWhitelist struct{}"
	}

	return strings.Join([]string{"ObjWhitelist", string(data)}, " ")
}
