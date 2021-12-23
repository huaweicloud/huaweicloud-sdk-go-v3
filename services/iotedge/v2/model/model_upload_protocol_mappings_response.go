package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadProtocolMappingsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadProtocolMappingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadProtocolMappingsResponse struct{}"
	}

	return strings.Join([]string{"UploadProtocolMappingsResponse", string(data)}, " ")
}
