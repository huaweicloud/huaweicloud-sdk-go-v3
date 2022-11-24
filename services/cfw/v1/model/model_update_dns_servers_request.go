package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDnsServersRequest struct {
	Body *UpdateDnsServersRequestBody `json:"body,omitempty"`
}

func (o UpdateDnsServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersRequest struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersRequest", string(data)}, " ")
}
