package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterV2TenantImageRequest Request Object
type RegisterV2TenantImageRequest struct {
	Body *RegisterV2ImageRequestInput `json:"body,omitempty"`
}

func (o RegisterV2TenantImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterV2TenantImageRequest struct{}"
	}

	return strings.Join([]string{"RegisterV2TenantImageRequest", string(data)}, " ")
}
