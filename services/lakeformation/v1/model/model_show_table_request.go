package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableRequest Request Object
type ShowTableRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// catalog名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	CatalogName string `json:"catalog_name"`

	// 数据库名称。只能包含中文、字母、数字、下划线、中划线，且长度为1~128个字符。
	DatabaseName string `json:"database_name"`

	// 表名称。只能包含中文、字母、数字、下划线、中划线，且长度为1~256个字符。
	TableName string `json:"table_name"`

	// 版本ID，默认为最新版本
	VersionId *string `json:"version_id,omitempty"`
}

func (o ShowTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableRequest struct{}"
	}

	return strings.Join([]string{"ShowTableRequest", string(data)}, " ")
}
