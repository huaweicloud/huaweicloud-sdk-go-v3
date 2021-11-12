package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckBackendConnectivityRequest struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *CheckBackendConnectivityReq `json:"body,omitempty"`
}

func (o CheckBackendConnectivityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckBackendConnectivityRequest struct{}"
	}

	return strings.Join([]string{"CheckBackendConnectivityRequest", string(data)}, " ")
}
