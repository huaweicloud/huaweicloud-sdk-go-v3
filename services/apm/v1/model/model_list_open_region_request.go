package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOpenRegionRequest Request Object
type ListOpenRegionRequest struct {
}

func (o ListOpenRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpenRegionRequest struct{}"
	}

	return strings.Join([]string{"ListOpenRegionRequest", string(data)}, " ")
}
