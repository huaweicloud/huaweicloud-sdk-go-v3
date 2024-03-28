package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableRequest Request Object
type ShowTableRequest struct {

	// 待描述的表所在的数据库名称。
	DatabaseName string `json:"database_name"`

	// 待描述表的名称。
	TableName string `json:"table_name"`
}

func (o ShowTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableRequest struct{}"
	}

	return strings.Join([]string{"ShowTableRequest", string(data)}, " ")
}
