package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateModuleMetadataResponse Response Object
type UpdatePrivateModuleMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePrivateModuleMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateModuleMetadataResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateModuleMetadataResponse", string(data)}, " ")
}
