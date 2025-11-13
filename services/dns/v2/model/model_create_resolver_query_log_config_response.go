package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResolverQueryLogConfigResponse Response Object
type CreateResolverQueryLogConfigResponse struct {
	ResolverQueryLogConfig *ResolverQueryLogConfig `json:"resolver_query_log_config,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o CreateResolverQueryLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolverQueryLogConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateResolverQueryLogConfigResponse", string(data)}, " ")
}
