package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsPolicy obs桶存放的策略。
type ObsPolicy struct {

	// 版本号。
	Version *string `json:"version,omitempty"`

	Statement *ObsPolicyStatement `json:"statement,omitempty"`
}

func (o ObsPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsPolicy struct{}"
	}

	return strings.Join([]string{"ObsPolicy", string(data)}, " ")
}
