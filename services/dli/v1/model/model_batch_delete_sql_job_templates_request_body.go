package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteSqlJobTemplatesRequestBody struct {

	// 待删除的sql模板ID列表。
	SqlIds []string `json:"sql_ids"`
}

func (o BatchDeleteSqlJobTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSqlJobTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteSqlJobTemplatesRequestBody", string(data)}, " ")
}
