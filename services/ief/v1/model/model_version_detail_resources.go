package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源配额
type VersionDetailResources struct {
	Limits *VersionDetailResourcesLimits `json:"limits,omitempty"`

	Requests *VersionDetailResourcesRequests `json:"requests,omitempty"`
}

func (o VersionDetailResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDetailResources struct{}"
	}

	return strings.Join([]string{"VersionDetailResources", string(data)}, " ")
}
