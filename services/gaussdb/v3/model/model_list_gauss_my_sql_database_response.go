package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlDatabaseResponse Response Object
type ListGaussMySqlDatabaseResponse struct {

	// 数据库信息列表。
	Databases *[]ListGaussMysqlDatabaseInfo `json:"databases,omitempty"`

	// 数据库总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGaussMySqlDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlDatabaseResponse", string(data)}, " ")
}
