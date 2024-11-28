package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyDatabaseRequest Request Object
type CopyDatabaseRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *AddCopyDatabaseRequestBody `json:"body,omitempty"`
}

func (o CopyDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CopyDatabaseRequest", string(data)}, " ")
}
