package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostHeaderInfo struct {

	// 主机
	Host *string `json:"Host,omitempty"`
}

func (o HostHeaderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostHeaderInfo struct{}"
	}

	return strings.Join([]string{"HostHeaderInfo", string(data)}, " ")
}
