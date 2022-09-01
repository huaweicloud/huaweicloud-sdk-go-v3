package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetGaussMySqlDatabasePasswordResponse struct {

	// 修改数据库用户密码的任务id。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetGaussMySqlDatabasePasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetGaussMySqlDatabasePasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetGaussMySqlDatabasePasswordResponse", string(data)}, " ")
}
