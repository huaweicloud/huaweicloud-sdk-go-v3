package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NovaAttachInterfaceRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *NovaAttachInterfaceRequestBody `json:"body,omitempty"`
}

func (o NovaAttachInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceRequest struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceRequest", string(data)}, " ")
}
