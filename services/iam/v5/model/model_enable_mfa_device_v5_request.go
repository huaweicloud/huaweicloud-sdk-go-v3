package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableMfaDeviceV5Request Request Object
type EnableMfaDeviceV5Request struct {
	Body *EnableMfaDeviceReqBody `json:"body,omitempty"`
}

func (o EnableMfaDeviceV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableMfaDeviceV5Request struct{}"
	}

	return strings.Join([]string{"EnableMfaDeviceV5Request", string(data)}, " ")
}
