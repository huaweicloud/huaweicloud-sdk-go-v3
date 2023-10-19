package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CoveredIpvo struct {

	// ip地址
	Ip *string `json:"ip,omitempty"`

	// 覆盖ip地址。
	CoveredIp *string `json:"covered_Ip,omitempty"`
}

func (o CoveredIpvo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CoveredIpvo struct{}"
	}

	return strings.Join([]string{"CoveredIpvo", string(data)}, " ")
}
