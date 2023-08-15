package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGaussMySqlBackupResponse Response Object
type DeleteGaussMySqlBackupResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlBackupResponse", string(data)}, " ")
}
