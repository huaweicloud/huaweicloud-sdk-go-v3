package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateGaussMySqlDatabaseUserCommentResponse struct {

	// 修改数据库用户备注的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlDatabaseUserCommentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlDatabaseUserCommentResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlDatabaseUserCommentResponse", string(data)}, " ")
}
