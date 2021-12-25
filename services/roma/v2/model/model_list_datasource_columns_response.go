package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDatasourceColumnsResponse struct {
	// 返回的实体对象

	Columns        *[]ColumnInfo `json:"columns,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDatasourceColumnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourceColumnsResponse struct{}"
	}

	return strings.Join([]string{"ListDatasourceColumnsResponse", string(data)}, " ")
}
