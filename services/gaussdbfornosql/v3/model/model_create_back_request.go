package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateBackRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *NoSqlCreateBackupRequestBody `json:"body,omitempty"`
}

func (o CreateBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackRequest struct{}"
	}

	return strings.Join([]string{"CreateBackRequest", string(data)}, " ")
}
