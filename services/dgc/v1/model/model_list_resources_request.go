package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResourcesRequest struct {
}

func (o ListResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}
