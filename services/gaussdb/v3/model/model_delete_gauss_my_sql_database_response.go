package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlDatabaseResponse Response Object
type DeleteGaussMySqlDatabaseResponse struct {

	// 删除数据库的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlDatabaseResponse", string(data)}, " ")
}
