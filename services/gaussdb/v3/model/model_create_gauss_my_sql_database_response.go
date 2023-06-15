package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateGaussMySqlDatabaseResponse struct {

	// 创建数据库的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMySqlDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlDatabaseResponse", string(data)}, " ")
}
