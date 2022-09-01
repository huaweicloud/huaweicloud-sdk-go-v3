package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CompareTableInfoWithToken struct {

	// 表名。
	TableName string `json:"table_name" xml:"table_name"`

	// 该表的min token。
	MinToken *string `json:"min_token,omitempty" xml:"min_token"`

	// 该表的max token。
	MaxToken *string `json:"max_token,omitempty" xml:"max_token"`
}

func (o CompareTableInfoWithToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTableInfoWithToken struct{}"
	}

	return strings.Join([]string{"CompareTableInfoWithToken", string(data)}, " ")
}
