package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncDnInformationRequest Request Object
type SyncDnInformationRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`
}

func (o SyncDnInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncDnInformationRequest struct{}"
	}

	return strings.Join([]string{"SyncDnInformationRequest", string(data)}, " ")
}
