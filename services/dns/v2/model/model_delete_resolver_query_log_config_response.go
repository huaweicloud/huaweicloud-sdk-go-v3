package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResolverQueryLogConfigResponse Response Object
type DeleteResolverQueryLogConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResolverQueryLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResolverQueryLogConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteResolverQueryLogConfigResponse", string(data)}, " ")
}
