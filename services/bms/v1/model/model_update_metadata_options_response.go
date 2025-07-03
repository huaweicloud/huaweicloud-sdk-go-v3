package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMetadataOptionsResponse Response Object
type UpdateMetadataOptionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMetadataOptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMetadataOptionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateMetadataOptionsResponse", string(data)}, " ")
}
