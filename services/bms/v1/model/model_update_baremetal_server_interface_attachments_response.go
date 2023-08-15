package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBaremetalServerInterfaceAttachmentsResponse Response Object
type UpdateBaremetalServerInterfaceAttachmentsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateBaremetalServerInterfaceAttachmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerInterfaceAttachmentsResponse struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerInterfaceAttachmentsResponse", string(data)}, " ")
}
