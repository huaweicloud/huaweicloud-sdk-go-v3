package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseRequest Request Object
type CreateDatabaseRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如:2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// catalog名称。只能包含字母、数字和下划线,且长度为1~256个字符。
	CatalogName string `json:"catalog_name"`

	Body *DatabaseInput `json:"body,omitempty"`
}

func (o CreateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequest", string(data)}, " ")
}
