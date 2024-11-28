package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddCopyDatabaseRequestBody struct {

	// 操作名称(copy_database)
	ProcedureName *string `json:"procedure_name,omitempty"`

	// 源库和目的库信息
	Params *interface{} `json:"params,omitempty"`
}

func (o AddCopyDatabaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCopyDatabaseRequestBody struct{}"
	}

	return strings.Join([]string{"AddCopyDatabaseRequestBody", string(data)}, " ")
}
