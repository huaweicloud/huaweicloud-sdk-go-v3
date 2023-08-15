package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabasesList 数据库信息
type DatabasesList struct {

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库描述
	Description *string `json:"description,omitempty"`
}

func (o DatabasesList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabasesList struct{}"
	}

	return strings.Join([]string{"DatabasesList", string(data)}, " ")
}
