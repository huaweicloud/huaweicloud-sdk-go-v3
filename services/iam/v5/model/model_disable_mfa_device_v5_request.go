package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableMfaDeviceV5Request Request Object
type DisableMfaDeviceV5Request struct {
	Body *DisableMfaDeviceReqBody `json:"body,omitempty"`
}

func (o DisableMfaDeviceV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableMfaDeviceV5Request struct{}"
	}

	return strings.Join([]string{"DisableMfaDeviceV5Request", string(data)}, " ")
}
