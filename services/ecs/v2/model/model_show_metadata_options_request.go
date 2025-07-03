package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetadataOptionsRequest Request Object
type ShowMetadataOptionsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowMetadataOptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadataOptionsRequest struct{}"
	}

	return strings.Join([]string{"ShowMetadataOptionsRequest", string(data)}, " ")
}
