package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlTemplateDatabasesResponse Response Object
type ListSqlTemplateDatabasesResponse struct {

	// 数据库列表
	DbNameList     *[]string `json:"db_name_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSqlTemplateDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTemplateDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlTemplateDatabasesResponse", string(data)}, " ")
}
