package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostStatus 服务器状态
type HostStatus struct {
}

func (o HostStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostStatus struct{}"
	}

	return strings.Join([]string{"HostStatus", string(data)}, " ")
}
