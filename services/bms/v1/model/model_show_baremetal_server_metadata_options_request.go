package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaremetalServerMetadataOptionsRequest Request Object
type ShowBaremetalServerMetadataOptionsRequest struct {

	// 裸金属服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowBaremetalServerMetadataOptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerMetadataOptionsRequest struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerMetadataOptionsRequest", string(data)}, " ")
}
