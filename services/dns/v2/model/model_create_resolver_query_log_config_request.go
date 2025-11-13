package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResolverQueryLogConfigRequest Request Object
type CreateResolverQueryLogConfigRequest struct {
	Body *CreateResolverQueryLogConfigRequestBody `json:"body,omitempty"`
}

func (o CreateResolverQueryLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolverQueryLogConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateResolverQueryLogConfigRequest", string(data)}, " ")
}
