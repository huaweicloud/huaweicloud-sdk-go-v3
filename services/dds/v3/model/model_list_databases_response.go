package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesResponse Response Object
type ListDatabasesResponse struct {

	// 数据库列表信息。
	Databases *[]DatabaseInfoItem `json:"databases,omitempty"`

	// 数据库列表总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabasesResponse", string(data)}, " ")
}
