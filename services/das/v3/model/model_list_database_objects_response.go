package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseObjectsResponse Response Object
type ListDatabaseObjectsResponse struct {

	// 数据库对象信息列表
	Data *interface{} `json:"data,omitempty"`

	// 列表大小
	Total *int64 `json:"total,omitempty"`

	// 对象类型
	ObjectType     *string `json:"object_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDatabaseObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseObjectsResponse", string(data)}, " ")
}
