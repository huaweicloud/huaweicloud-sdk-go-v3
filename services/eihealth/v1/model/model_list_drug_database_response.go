package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDrugDatabaseResponse Response Object
type ListDrugDatabaseResponse struct {

	// 数据库列表
	Databases *[]DrugDatabaseDto `json:"databases,omitempty"`

	// 数据库总数
	Count *int32 `json:"count,omitempty"`

	// 当前用户数据库总数
	CurUserCount   *int32 `json:"cur_user_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDrugDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrugDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ListDrugDatabaseResponse", string(data)}, " ")
}
