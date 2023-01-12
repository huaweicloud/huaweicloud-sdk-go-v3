package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateBackup2Response struct {

	// 图备份任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBackup2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackup2Response struct{}"
	}

	return strings.Join([]string{"CreateBackup2Response", string(data)}, " ")
}
