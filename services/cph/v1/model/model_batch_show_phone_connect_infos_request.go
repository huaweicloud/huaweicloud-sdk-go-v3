package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowPhoneConnectInfosRequest Request Object
type BatchShowPhoneConnectInfosRequest struct {
	Body *ConnectionRequestBody `json:"body,omitempty"`
}

func (o BatchShowPhoneConnectInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPhoneConnectInfosRequest struct{}"
	}

	return strings.Join([]string{"BatchShowPhoneConnectInfosRequest", string(data)}, " ")
}
