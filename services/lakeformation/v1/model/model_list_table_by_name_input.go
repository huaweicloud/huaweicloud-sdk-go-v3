package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableByNameInput 表列表
type ListTableByNameInput struct {

	// 被查询表的名字列表
	TableNames []string `json:"table_names"`
}

func (o ListTableByNameInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableByNameInput struct{}"
	}

	return strings.Join([]string{"ListTableByNameInput", string(data)}, " ")
}
