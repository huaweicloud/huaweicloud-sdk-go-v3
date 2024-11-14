package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbCacheMappingRequest Request Object
type CreateDbCacheMappingRequest struct {
	Body *CreateDbCacheMappingRequestBody `json:"body,omitempty"`
}

func (o CreateDbCacheMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbCacheMappingRequest struct{}"
	}

	return strings.Join([]string{"CreateDbCacheMappingRequest", string(data)}, " ")
}
