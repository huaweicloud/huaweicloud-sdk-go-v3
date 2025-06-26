package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRegistryResponse Response Object
type UpdateInstanceRegistryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRegistryResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRegistryResponse", string(data)}, " ")
}
