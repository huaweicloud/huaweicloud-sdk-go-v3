package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHomeRegionRequest Request Object
type ShowHomeRegionRequest struct {
}

func (o ShowHomeRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHomeRegionRequest struct{}"
	}

	return strings.Join([]string{"ShowHomeRegionRequest", string(data)}, " ")
}
