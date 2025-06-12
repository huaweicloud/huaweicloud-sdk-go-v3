package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerMetadataOptionsRequest Request Object
type UpdateServerMetadataOptionsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *UpdateServerMetadataOptionsRequestBody `json:"body,omitempty"`
}

func (o UpdateServerMetadataOptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerMetadataOptionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerMetadataOptionsRequest", string(data)}, " ")
}
