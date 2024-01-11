package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaShowServerInterfaceResponse Response Object
type NovaShowServerInterfaceResponse struct {
	InterfaceAttachment *NovaServerInterfaceDetail `json:"interfaceAttachment,omitempty"`
	HttpStatusCode      int                        `json:"-"`
}

func (o NovaShowServerInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowServerInterfaceResponse struct{}"
	}

	return strings.Join([]string{"NovaShowServerInterfaceResponse", string(data)}, " ")
}
