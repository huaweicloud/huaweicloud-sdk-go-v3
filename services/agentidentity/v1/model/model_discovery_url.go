package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiscoveryUrl This URL is used to fetch OpenID Connect configuration.
type DiscoveryUrl struct {
}

func (o DiscoveryUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscoveryUrl struct{}"
	}

	return strings.Join([]string{"DiscoveryUrl", string(data)}, " ")
}
