package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerMetadataResponse Response Object
type DeleteServerMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerMetadataResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerMetadataResponse", string(data)}, " ")
}
