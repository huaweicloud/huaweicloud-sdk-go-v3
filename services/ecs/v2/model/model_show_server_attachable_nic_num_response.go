package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerAttachableNicNumResponse Response Object
type ShowServerAttachableNicNumResponse struct {
	AttachableQuantity *AttachableQuantityForNic `json:"attachableQuantity,omitempty"`

	InterfaceAttachments *[]InterfaceExt `json:"interfaceAttachments,omitempty"`
	HttpStatusCode       int             `json:"-"`
}

func (o ShowServerAttachableNicNumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerAttachableNicNumResponse struct{}"
	}

	return strings.Join([]string{"ShowServerAttachableNicNumResponse", string(data)}, " ")
}
