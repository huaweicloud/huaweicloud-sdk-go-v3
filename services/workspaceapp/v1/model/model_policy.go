package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Policy OBS桶存放的策略。
type Policy struct {

	// 版本号。
	Version *string `json:"version,omitempty"`

	Statement *ObsPolicyStatement `json:"statement,omitempty"`
}

func (o Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policy struct{}"
	}

	return strings.Join([]string{"Policy", string(data)}, " ")
}
