package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaAttachInterfaceResponse Response Object
type NovaAttachInterfaceResponse struct {
	InterfaceAttachment *NovaServerInterfaceDetail `json:"interfaceAttachment,omitempty"`
	HttpStatusCode      int                        `json:"-"`
}

func (o NovaAttachInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceResponse struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceResponse", string(data)}, " ")
}
