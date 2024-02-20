package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanClientsRequest Request Object
type ScanClientsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ScanClientsRequestBody `json:"body,omitempty"`
}

func (o ScanClientsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanClientsRequest struct{}"
	}

	return strings.Join([]string{"ScanClientsRequest", string(data)}, " ")
}
