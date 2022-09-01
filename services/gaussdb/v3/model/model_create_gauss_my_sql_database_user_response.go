package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateGaussMySqlDatabaseUserResponse struct {

	// 创建数据库用户的任务id。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMySqlDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlDatabaseUserResponse", string(data)}, " ")
}
