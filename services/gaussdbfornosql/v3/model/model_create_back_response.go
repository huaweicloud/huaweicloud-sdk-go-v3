package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateBackResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 备份ID。
	BackupId       *string `json:"backup_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackResponse struct{}"
	}

	return strings.Join([]string{"CreateBackResponse", string(data)}, " ")
}
