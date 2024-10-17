package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRdsDatabasesResponse Response Object
type ListRdsDatabasesResponse struct {

	// rds数据库列表
	Databases *[]RdsDbListResponseDatabases `json:"databases,omitempty"`

	// 总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRdsDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRdsDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListRdsDatabasesResponse", string(data)}, " ")
}
