package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NonRequiredHostedCloud 归属云。
type NonRequiredHostedCloud struct {
	HostedCloud *HostedCloudEnum `json:"hosted_cloud,omitempty"`
}

func (o NonRequiredHostedCloud) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonRequiredHostedCloud struct{}"
	}

	return strings.Join([]string{"NonRequiredHostedCloud", string(data)}, " ")
}
