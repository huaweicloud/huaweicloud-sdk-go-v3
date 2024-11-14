package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbCacheMappingRequest Request Object
type DeleteDbCacheMappingRequest struct {
	Body *DeleteDbCacheMappingRequestBody `json:"body,omitempty"`
}

func (o DeleteDbCacheMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbCacheMappingRequest struct{}"
	}

	return strings.Join([]string{"DeleteDbCacheMappingRequest", string(data)}, " ")
}
