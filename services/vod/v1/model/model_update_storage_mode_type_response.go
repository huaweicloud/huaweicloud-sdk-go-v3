package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStorageModeTypeResponse Response Object
type UpdateStorageModeTypeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateStorageModeTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStorageModeTypeResponse struct{}"
	}

	return strings.Join([]string{"UpdateStorageModeTypeResponse", string(data)}, " ")
}
