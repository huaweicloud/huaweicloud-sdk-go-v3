package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeadLockDatabasesResponse Response Object
type ListDeadLockDatabasesResponse struct {

	// 数据库名列表
	DbList         *[]string `json:"db_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDeadLockDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeadLockDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListDeadLockDatabasesResponse", string(data)}, " ")
}
