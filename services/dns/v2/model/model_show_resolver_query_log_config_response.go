package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResolverQueryLogConfigResponse Response Object
type ShowResolverQueryLogConfigResponse struct {
	ResolverQueryLogConfig *ResolverQueryLogConfig `json:"resolver_query_log_config,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o ShowResolverQueryLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResolverQueryLogConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowResolverQueryLogConfigResponse", string(data)}, " ")
}
