package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CompareObjectInfoWithToken struct {

	// 库名。
	DbName string `json:"db_name" xml:"db_name"`

	// 该库下的表信息列表（带token）。
	TableNameWithToken *[]CompareTableInfoWithToken `json:"table_name_with_token,omitempty" xml:"table_name_with_token"`
}

func (o CompareObjectInfoWithToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareObjectInfoWithToken struct{}"
	}

	return strings.Join([]string{"CompareObjectInfoWithToken", string(data)}, " ")
}
