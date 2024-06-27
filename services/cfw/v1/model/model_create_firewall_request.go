package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFirewallRequest Request Object
type CreateFirewallRequest struct {

	// 保证客户端请求幂等性的标识。  该标识为32位UUID格式，由客户端生成，且需确保不同请求之间该标识具有唯一性。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateFirewallReq `json:"body,omitempty"`
}

func (o CreateFirewallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallRequest struct{}"
	}

	return strings.Join([]string{"CreateFirewallRequest", string(data)}, " ")
}
