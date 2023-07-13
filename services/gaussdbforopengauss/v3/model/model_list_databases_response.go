package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesResponse Response Object
type ListDatabasesResponse struct {

	// 列表中每个元素表示一个数据库。
	Databases *[]GaussDBforOpenGaussListDatabase `json:"databases,omitempty"`

	// 数据库总数。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabasesResponse", string(data)}, " ")
}
