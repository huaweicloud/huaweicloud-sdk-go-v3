package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateListenerRequest Request Object
type DisassociateListenerRequest struct {

	// IP地址组ID。
	IpGroupId string `json:"ip_group_id"`

	Body *DisassociateListenerRequestBody `json:"body,omitempty"`
}

func (o DisassociateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateListenerRequest struct{}"
	}

	return strings.Join([]string{"DisassociateListenerRequest", string(data)}, " ")
}
