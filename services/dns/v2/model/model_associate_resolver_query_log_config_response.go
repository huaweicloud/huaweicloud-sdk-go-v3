package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateResolverQueryLogConfigResponse Response Object
type AssociateResolverQueryLogConfigResponse struct {
	ResolverQueryLogConfig *ResolverQueryLogConfig `json:"resolver_query_log_config,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o AssociateResolverQueryLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResolverQueryLogConfigResponse struct{}"
	}

	return strings.Join([]string{"AssociateResolverQueryLogConfigResponse", string(data)}, " ")
}
