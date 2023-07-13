package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatasourceTablesResponse Response Object
type ListDatasourceTablesResponse struct {

	// 数据源中所有的表名称
	Tables         *[]string `json:"tables,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDatasourceTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourceTablesResponse struct{}"
	}

	return strings.Join([]string{"ListDatasourceTablesResponse", string(data)}, " ")
}
