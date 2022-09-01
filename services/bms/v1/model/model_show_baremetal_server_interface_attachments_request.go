package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBaremetalServerInterfaceAttachmentsRequest struct {

	// 裸金属服务器ID
	ServerId string `json:"server_id" xml:"server_id"`
}

func (o ShowBaremetalServerInterfaceAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerInterfaceAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerInterfaceAttachmentsRequest", string(data)}, " ")
}
