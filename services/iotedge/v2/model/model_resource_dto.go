package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceDto struct {
	Limits *ResourceConfigDto `json:"limits,omitempty" xml:"limits"`

	Requests *ResourceConfigDto `json:"requests,omitempty" xml:"requests"`
}

func (o ResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}
