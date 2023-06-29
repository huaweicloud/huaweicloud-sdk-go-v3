package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesResponse Response Object
type ListDatabasesResponse struct {

	// 当前数据连接数据库记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 数据库列表
	Databases      *[]DatabasesList `json:"databases,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabasesResponse", string(data)}, " ")
}
