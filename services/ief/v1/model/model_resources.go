package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容器使用的资源
type Resources struct {
	Limits *LimitsRequests `json:"limits,omitempty" xml:"limits"`

	Requests *LimitsRequests `json:"requests,omitempty" xml:"requests"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
