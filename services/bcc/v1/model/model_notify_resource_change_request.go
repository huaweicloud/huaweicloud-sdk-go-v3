package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotifyResourceChangeRequest Request Object
type NotifyResourceChangeRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *ResourceNotifyBody `json:"body,omitempty"`
}

func (o NotifyResourceChangeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifyResourceChangeRequest struct{}"
	}

	return strings.Join([]string{"NotifyResourceChangeRequest", string(data)}, " ")
}
