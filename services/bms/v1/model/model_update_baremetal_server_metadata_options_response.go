package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBaremetalServerMetadataOptionsResponse Response Object
type UpdateBaremetalServerMetadataOptionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateBaremetalServerMetadataOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerMetadataOptionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerMetadataOptionsResponse", string(data)}, " ")
}
