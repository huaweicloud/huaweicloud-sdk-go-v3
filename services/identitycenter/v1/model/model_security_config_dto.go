package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigDto 证书安全配置
type SecurityConfigDto struct {

	// 证书过期时间
	Ttl string `json:"ttl"`
}

func (o SecurityConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigDto struct{}"
	}

	return strings.Join([]string{"SecurityConfigDto", string(data)}, " ")
}
