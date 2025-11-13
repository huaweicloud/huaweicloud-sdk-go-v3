package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateResolverQueryLogConfigRequest Request Object
type AssociateResolverQueryLogConfigRequest struct {

	// 访问日志ID。
	Id string `json:"id"`

	Body *AssociateResolverQueryLogConfigRequestBody `json:"body,omitempty"`
}

func (o AssociateResolverQueryLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResolverQueryLogConfigRequest struct{}"
	}

	return strings.Join([]string{"AssociateResolverQueryLogConfigRequest", string(data)}, " ")
}
