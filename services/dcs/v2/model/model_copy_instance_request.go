package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CopyInstanceRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *BackupInstanceBody `json:"body,omitempty"`
}

func (o CopyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyInstanceRequest struct{}"
	}

	return strings.Join([]string{"CopyInstanceRequest", string(data)}, " ")
}
