package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataProfileRequest Request Object
type ShowDataProfileRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 数据连接ID
	DwId string `json:"dw_id"`

	// 数据库类型
	DbType string `json:"db_type"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// 表名
	TableName string `json:"table_name"`
}

func (o ShowDataProfileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataProfileRequest struct{}"
	}

	return strings.Join([]string{"ShowDataProfileRequest", string(data)}, " ")
}
