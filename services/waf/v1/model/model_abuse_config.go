package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AbuseConfig 算力滥用防护配置
type AbuseConfig struct {
	Scope *AbuseScope `json:"scope,omitempty"`

	Threshold *AbuseThreshold `json:"threshold,omitempty"`
}

func (o AbuseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AbuseConfig struct{}"
	}

	return strings.Join([]string{"AbuseConfig", string(data)}, " ")
}
