package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDrugDatabaseFileRequest Request Object
type AddDrugDatabaseFileRequest struct {

	// 数据库id
	DatabaseId string `json:"database_id"`

	Body *AddDrugDatabaseFileReq `json:"body,omitempty"`
}

func (o AddDrugDatabaseFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDrugDatabaseFileRequest struct{}"
	}

	return strings.Join([]string{"AddDrugDatabaseFileRequest", string(data)}, " ")
}
