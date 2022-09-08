package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRestoreDatabasesResponse struct {

	// 数据库总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据库列表，列表中每个元素表示一个数据库。
	Databases      *[]string `json:"databases,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRestoreDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreDatabasesResponse", string(data)}, " ")
}
