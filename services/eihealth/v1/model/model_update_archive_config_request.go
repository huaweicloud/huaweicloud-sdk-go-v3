package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateArchiveConfigRequest Request Object
type UpdateArchiveConfigRequest struct {

	// region id
	RegionId string `json:"region_id"`
}

func (o UpdateArchiveConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateArchiveConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateArchiveConfigRequest", string(data)}, " ")
}
