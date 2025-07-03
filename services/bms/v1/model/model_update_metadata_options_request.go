package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMetadataOptionsRequest Request Object
type UpdateMetadataOptionsRequest struct {

	// 裸金属服务器ID。
	ServerId string `json:"server_id"`

	Body *UpdateBaremetalServerMetadataOptionsRequestBody `json:"body,omitempty"`
}

func (o UpdateMetadataOptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMetadataOptionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateMetadataOptionsRequest", string(data)}, " ")
}
