package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateListenerRequest Request Object
type AssociateListenerRequest struct {

	// IP地址组ID。
	IpGroupId string `json:"ip_group_id"`

	Body *AssociateListenerRequestBody `json:"body,omitempty"`
}

func (o AssociateListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateListenerRequest struct{}"
	}

	return strings.Join([]string{"AssociateListenerRequest", string(data)}, " ")
}
