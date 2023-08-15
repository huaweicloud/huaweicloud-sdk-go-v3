package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGaussMySqlDatabaseUserResponse Response Object
type CreateGaussMySqlDatabaseUserResponse struct {

	// 创建数据库用户的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMySqlDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlDatabaseUserResponse", string(data)}, " ")
}
