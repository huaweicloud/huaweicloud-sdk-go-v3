package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerMetadataOptionsResponse Response Object
type UpdateServerMetadataOptionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateServerMetadataOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerMetadataOptionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerMetadataOptionsResponse", string(data)}, " ")
}
