package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateResolverQueryLogConfigRequest Request Object
type DisassociateResolverQueryLogConfigRequest struct {

	// 访问日志ID。
	Id string `json:"id"`

	Body *AssociateResolverQueryLogConfigRequestBody `json:"body,omitempty"`
}

func (o DisassociateResolverQueryLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResolverQueryLogConfigRequest struct{}"
	}

	return strings.Join([]string{"DisassociateResolverQueryLogConfigRequest", string(data)}, " ")
}
