package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableMfaDeviceV5Response Response Object
type DisableMfaDeviceV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableMfaDeviceV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableMfaDeviceV5Response struct{}"
	}

	return strings.Join([]string{"DisableMfaDeviceV5Response", string(data)}, " ")
}
