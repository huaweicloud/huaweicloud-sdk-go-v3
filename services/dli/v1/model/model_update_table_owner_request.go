package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTableOwnerRequest Request Object
type UpdateTableOwnerRequest struct {

	// 删除的数据库名称。
	DatabaseName string `json:"database_name"`

	TableName string `json:"table_name"`

	Body *UpdateOwnerRequestBody `json:"body,omitempty"`
}

func (o UpdateTableOwnerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableOwnerRequest struct{}"
	}

	return strings.Join([]string{"UpdateTableOwnerRequest", string(data)}, " ")
}
