package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceDto struct {
	Limits *ResourceConfigDto `json:"limits,omitempty"`

	Requests *ResourceConfigDto `json:"requests,omitempty"`
}

func (o ResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}
