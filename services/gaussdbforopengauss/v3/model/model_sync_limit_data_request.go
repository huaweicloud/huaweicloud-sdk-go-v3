package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncLimitDataRequest Request Object
type SyncLimitDataRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o SyncLimitDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncLimitDataRequest struct{}"
	}

	return strings.Join([]string{"SyncLimitDataRequest", string(data)}, " ")
}
