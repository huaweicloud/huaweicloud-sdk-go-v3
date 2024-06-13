package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallTagsRequest Request Object
type ListFirewallTagsRequest struct {
}

func (o ListFirewallTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallTagsRequest struct{}"
	}

	return strings.Join([]string{"ListFirewallTagsRequest", string(data)}, " ")
}
