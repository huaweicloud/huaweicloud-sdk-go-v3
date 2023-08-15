package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBaremetalServerInterfaceAttachmentsRequest Request Object
type UpdateBaremetalServerInterfaceAttachmentsRequest struct {
	PortId string `json:"port_id"`

	ServerId string `json:"server_id"`

	Body *UpdateBaremetalServerInterfaceAttachmentsReq `json:"body,omitempty"`
}

func (o UpdateBaremetalServerInterfaceAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerInterfaceAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerInterfaceAttachmentsRequest", string(data)}, " ")
}
