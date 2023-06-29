package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaAttachInterfaceRequestBody This is a auto create Body Object
type NovaAttachInterfaceRequestBody struct {
	InterfaceAttachment *NovaAttachInterfaceOption `json:"interfaceAttachment"`
}

func (o NovaAttachInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceRequestBody", string(data)}, " ")
}
