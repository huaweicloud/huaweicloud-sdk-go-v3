package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBaremetalServerMetadataOptionsRequest Request Object
type UpdateBaremetalServerMetadataOptionsRequest struct {

	// 裸金属服务器ID。
	ServerId string `json:"server_id"`

	Body *UpdateBaremetalServerMetadataOptionsRequestBody `json:"body,omitempty"`
}

func (o UpdateBaremetalServerMetadataOptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerMetadataOptionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerMetadataOptionsRequest", string(data)}, " ")
}
