package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateGaussMySqlDatabaseCommentResponse struct {

	// 修改数据库备注的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlDatabaseCommentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlDatabaseCommentResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlDatabaseCommentResponse", string(data)}, " ")
}
