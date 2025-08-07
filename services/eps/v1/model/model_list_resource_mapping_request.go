package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceMappingRequest Request Object
type ListResourceMappingRequest struct {
}

func (o ListResourceMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceMappingRequest struct{}"
	}

	return strings.Join([]string{"ListResourceMappingRequest", string(data)}, " ")
}
