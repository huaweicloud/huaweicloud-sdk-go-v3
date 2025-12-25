package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapIpAddress IP地址
type IsapIpAddress struct {
}

func (o IsapIpAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapIpAddress struct{}"
	}

	return strings.Join([]string{"IsapIpAddress", string(data)}, " ")
}
