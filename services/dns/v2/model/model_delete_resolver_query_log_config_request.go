package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResolverQueryLogConfigRequest Request Object
type DeleteResolverQueryLogConfigRequest struct {

	// 访问日志ID。
	Id string `json:"id"`
}

func (o DeleteResolverQueryLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResolverQueryLogConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteResolverQueryLogConfigRequest", string(data)}, " ")
}
