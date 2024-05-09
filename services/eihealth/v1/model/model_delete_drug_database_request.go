package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugDatabaseRequest Request Object
type DeleteDrugDatabaseRequest struct {

	// 数据库id
	DatabaseId string `json:"database_id"`
}

func (o DeleteDrugDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteDrugDatabaseRequest", string(data)}, " ")
}
