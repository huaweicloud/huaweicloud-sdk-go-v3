package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateBaremetalServerMetadataResponse struct {
	Metadata       *KeyValue `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateBaremetalServerMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerMetadataResponse struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerMetadataResponse", string(data)}, " ")
}
