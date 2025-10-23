package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportedRegionRequest Request Object
type ListSupportedRegionRequest struct {
}

func (o ListSupportedRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedRegionRequest struct{}"
	}

	return strings.Join([]string{"ListSupportedRegionRequest", string(data)}, " ")
}
