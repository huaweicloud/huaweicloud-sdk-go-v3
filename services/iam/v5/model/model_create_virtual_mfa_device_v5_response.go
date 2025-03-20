package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirtualMfaDeviceV5Response Response Object
type CreateVirtualMfaDeviceV5Response struct {
	VirtualMfaDevice *VirtualMfaDevice `json:"virtual_mfa_device,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o CreateVirtualMfaDeviceV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualMfaDeviceV5Response struct{}"
	}

	return strings.Join([]string{"CreateVirtualMfaDeviceV5Response", string(data)}, " ")
}
