package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConnectorConnectivityRequest Request Object
type ValidateConnectorConnectivityRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SmartConnectValidateEntity `json:"body,omitempty"`
}

func (o ValidateConnectorConnectivityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConnectorConnectivityRequest struct{}"
	}

	return strings.Join([]string{"ValidateConnectorConnectivityRequest", string(data)}, " ")
}
