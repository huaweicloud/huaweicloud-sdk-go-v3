package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CompareObjectInfoWithToken struct {
	// 库名。

	DbName string `json:"db_name"`
	// 该库下的表信息列表（带token）。

	TableNameWithToken *[]CompareTableInfoWithToken `json:"table_name_with_token,omitempty"`
}

func (o CompareObjectInfoWithToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareObjectInfoWithToken struct{}"
	}

	return strings.Join([]string{"CompareObjectInfoWithToken", string(data)}, " ")
}
