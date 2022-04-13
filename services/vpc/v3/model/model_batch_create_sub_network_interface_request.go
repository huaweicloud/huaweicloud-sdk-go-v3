package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateSubNetworkInterfaceRequest struct {
	Body *BatchCreateSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSubNetworkInterfaceRequest", string(data)}, " ")
}
