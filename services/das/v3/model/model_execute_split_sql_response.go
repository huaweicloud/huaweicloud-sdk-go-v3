package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSplitSqlResponse Response Object
type ExecuteSplitSqlResponse struct {

	// 切分后的SQL语句列表
	SqlList        *[]string `json:"sql_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ExecuteSplitSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSplitSqlResponse struct{}"
	}

	return strings.Join([]string{"ExecuteSplitSqlResponse", string(data)}, " ")
}
