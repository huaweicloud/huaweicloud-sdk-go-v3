package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resources 容器使用的资源
type Resources struct {
	Limits *LimitsRequests `json:"limits,omitempty"`

	Requests *LimitsRequests `json:"requests,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
