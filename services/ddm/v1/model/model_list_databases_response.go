package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDatabasesResponse struct {
	// 逻辑库相关信息的集合

	Databases *[]GetDatabaseInfo `json:"databases,omitempty"`
	// 总条数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabasesResponse", string(data)}, " ")
}
