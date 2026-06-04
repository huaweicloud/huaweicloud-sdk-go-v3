package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableRequest Request Object
type CreateTableRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// catalog名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	CatalogName string `json:"catalog_name"`

	// 数据库名称。只能包含中文、字母、数字、下划线、中划线，且长度为1~128个字符。
	DatabaseName string `json:"database_name"`

	Body *CreateTableInput `json:"body,omitempty"`
}

func (o CreateTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableRequest struct{}"
	}

	return strings.Join([]string{"CreateTableRequest", string(data)}, " ")
}
