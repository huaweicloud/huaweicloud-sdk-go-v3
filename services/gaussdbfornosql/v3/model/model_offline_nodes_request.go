package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OfflineNodesRequest Request Object
type OfflineNodesRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OfflineNodesRequestBody `json:"body,omitempty"`
}

func (o OfflineNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfflineNodesRequest struct{}"
	}

	return strings.Join([]string{"OfflineNodesRequest", string(data)}, " ")
}
