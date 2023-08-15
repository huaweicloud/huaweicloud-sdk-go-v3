package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteSqlTemplatesRequestBody struct {

	// 待删除的sql模板ID列表。
	SqlIds []string `json:"sql_ids"`
}

func (o DeleteSqlTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteSqlTemplatesRequestBody", string(data)}, " ")
}
