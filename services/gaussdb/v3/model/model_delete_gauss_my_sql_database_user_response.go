package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteGaussMySqlDatabaseUserResponse struct {

	// 删除数据库用户的任务id。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlDatabaseUserResponse", string(data)}, " ")
}
