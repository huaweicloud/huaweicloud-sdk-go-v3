package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirtualMfaDeviceV5Request Request Object
type CreateVirtualMfaDeviceV5Request struct {
	Body *CreateVirtualMfaDeviceReqBody `json:"body,omitempty"`
}

func (o CreateVirtualMfaDeviceV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualMfaDeviceV5Request struct{}"
	}

	return strings.Join([]string{"CreateVirtualMfaDeviceV5Request", string(data)}, " ")
}
