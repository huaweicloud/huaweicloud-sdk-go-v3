package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SourceIpTopListInfoItems struct {

	// 攻击源ip地址
	Ip *string `json:"ip,omitempty"`

	// 攻击源ip攻击次数
	Num *int32 `json:"num,omitempty"`
}

func (o SourceIpTopListInfoItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceIpTopListInfoItems struct{}"
	}

	return strings.Join([]string{"SourceIpTopListInfoItems", string(data)}, " ")
}
