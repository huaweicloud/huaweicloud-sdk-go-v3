package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRedisPitrRequest Request Object
type RestoreRedisPitrRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RestoreRedisPitrRequestBody `json:"body,omitempty"`
}

func (o RestoreRedisPitrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRedisPitrRequest struct{}"
	}

	return strings.Join([]string{"RestoreRedisPitrRequest", string(data)}, " ")
}
