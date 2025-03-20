package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MfaEnabled 虚拟MFA设备是否开启。
type MfaEnabled struct {
}

func (o MfaEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaEnabled struct{}"
	}

	return strings.Join([]string{"MfaEnabled", string(data)}, " ")
}
