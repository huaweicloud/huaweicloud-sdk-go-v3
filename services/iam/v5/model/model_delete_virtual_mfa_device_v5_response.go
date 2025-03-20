package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVirtualMfaDeviceV5Response Response Object
type DeleteVirtualMfaDeviceV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVirtualMfaDeviceV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualMfaDeviceV5Response struct{}"
	}

	return strings.Join([]string{"DeleteVirtualMfaDeviceV5Response", string(data)}, " ")
}
