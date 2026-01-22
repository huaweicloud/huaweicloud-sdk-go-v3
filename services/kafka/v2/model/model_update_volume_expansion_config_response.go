package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVolumeExpansionConfigResponse Response Object
type UpdateVolumeExpansionConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateVolumeExpansionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeExpansionConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateVolumeExpansionConfigResponse", string(data)}, " ")
}
