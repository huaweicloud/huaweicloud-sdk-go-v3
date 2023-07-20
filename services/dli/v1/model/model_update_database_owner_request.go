package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseOwnerRequest Request Object
type UpdateDatabaseOwnerRequest struct {

	// 删除的数据库名称。
	DatabaseName string `json:"database_name"`

	Body *UpdateOwnerRequestBody `json:"body,omitempty"`
}

func (o UpdateDatabaseOwnerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseOwnerRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseOwnerRequest", string(data)}, " ")
}
