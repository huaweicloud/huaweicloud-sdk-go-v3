package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容器使用的资源
type Resources struct {
	Limits *ResourcesLimits `json:"limits,omitempty"`

	Requests *ResourcesRequests `json:"requests,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
