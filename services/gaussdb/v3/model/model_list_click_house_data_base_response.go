package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClickHouseDataBaseResponse Response Object
type ListClickHouseDataBaseResponse struct {

	// 数据库列表。
	Databases *[]string `json:"databases,omitempty"`

	// 数据库总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClickHouseDataBaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClickHouseDataBaseResponse struct{}"
	}

	return strings.Join([]string{"ListClickHouseDataBaseResponse", string(data)}, " ")
}
