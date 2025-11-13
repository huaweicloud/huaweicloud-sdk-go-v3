package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResolverQueryLogConfigRequest Request Object
type ShowResolverQueryLogConfigRequest struct {

	// 访问日志ID。
	Id string `json:"id"`
}

func (o ShowResolverQueryLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResolverQueryLogConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowResolverQueryLogConfigRequest", string(data)}, " ")
}
