package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MsdtcHostResult struct {

	// 主机名称
	HostName string `json:"host_name"`

	// 主机ip
	Ip string `json:"ip"`
}

func (o MsdtcHostResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MsdtcHostResult struct{}"
	}

	return strings.Join([]string{"MsdtcHostResult", string(data)}, " ")
}
