package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReplicationJobRequest Request Object
type DeleteReplicationJobRequest struct {

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 备份迁移任务ID。
	JobId string `json:"job_id"`
}

func (o DeleteReplicationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReplicationJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteReplicationJobRequest", string(data)}, " ")
}
