package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostNum 影响主机数量
type HostNum struct {
}

func (o HostNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostNum struct{}"
	}

	return strings.Join([]string{"HostNum", string(data)}, " ")
}
