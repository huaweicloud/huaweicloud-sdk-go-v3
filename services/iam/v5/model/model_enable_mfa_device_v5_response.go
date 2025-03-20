package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableMfaDeviceV5Response Response Object
type EnableMfaDeviceV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableMfaDeviceV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableMfaDeviceV5Response struct{}"
	}

	return strings.Join([]string{"EnableMfaDeviceV5Response", string(data)}, " ")
}
