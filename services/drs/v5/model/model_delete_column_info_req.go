package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteColumnInfoReq 删除对象列信息请求体
type DeleteColumnInfoReq struct {

	// 列所属表的id
	TableIds []string `json:"table_ids"`

	// 列所属schema的id
	SchemaIds []string `json:"schema_ids"`

	// 列所属库的id
	DbIds []string `json:"db_ids"`
}

func (o DeleteColumnInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteColumnInfoReq struct{}"
	}

	return strings.Join([]string{"DeleteColumnInfoReq", string(data)}, " ")
}
