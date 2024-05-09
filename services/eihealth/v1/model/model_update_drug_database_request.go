package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugDatabaseRequest Request Object
type UpdateDrugDatabaseRequest struct {

	// 数据库id
	DatabaseId string `json:"database_id"`

	Body *UpdateDrugDatabaseReq `json:"body,omitempty"`
}

func (o UpdateDrugDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugDatabaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateDrugDatabaseRequest", string(data)}, " ")
}
