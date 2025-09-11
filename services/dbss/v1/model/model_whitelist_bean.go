package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WhitelistBean struct {

	// 创建时间
	CreateTimestampMs *int64 `json:"create_timestamp_ms,omitempty"`

	// 数据库IDs
	DbIds *[]string `json:"db_ids,omitempty"`

	// 描述
	Desc *string `json:"desc,omitempty"`

	// 状态 - true:启用 - false: 禁用
	Enabled *bool `json:"enabled,omitempty"`

	// 记录ID
	Id *string `json:"id,omitempty"`

	// SQL语句
	Sql *string `json:"sql,omitempty"`

	// 更新时间
	UpdateTimestampMs *int64 `json:"update_timestamp_ms,omitempty"`
}

func (o WhitelistBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhitelistBean struct{}"
	}

	return strings.Join([]string{"WhitelistBean", string(data)}, " ")
}
