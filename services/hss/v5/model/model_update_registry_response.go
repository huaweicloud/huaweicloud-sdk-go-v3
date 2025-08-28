package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRegistryResponse Response Object
type UpdateRegistryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRegistryResponse struct{}"
	}

	return strings.Join([]string{"UpdateRegistryResponse", string(data)}, " ")
}
