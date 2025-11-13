package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateResolverQueryLogConfigResponse Response Object
type DisassociateResolverQueryLogConfigResponse struct {
	ResolverQueryLogConfig *ResolverQueryLogConfig `json:"resolver_query_log_config,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o DisassociateResolverQueryLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResolverQueryLogConfigResponse struct{}"
	}

	return strings.Join([]string{"DisassociateResolverQueryLogConfigResponse", string(data)}, " ")
}
