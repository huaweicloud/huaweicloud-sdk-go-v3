package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateHookMetadataResponse Response Object
type UpdatePrivateHookMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePrivateHookMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateHookMetadataResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateHookMetadataResponse", string(data)}, " ")
}
