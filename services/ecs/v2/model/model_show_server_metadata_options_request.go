package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerMetadataOptionsRequest Request Object
type ShowServerMetadataOptionsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowServerMetadataOptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerMetadataOptionsRequest struct{}"
	}

	return strings.Join([]string{"ShowServerMetadataOptionsRequest", string(data)}, " ")
}
