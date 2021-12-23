package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BrandBrokens struct {
	RecBytes *[]int32 `json:"recBytes,omitempty"`

	SentBytes *[]int32 `json:"sentBytes,omitempty"`
}

func (o BrandBrokens) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BrandBrokens struct{}"
	}

	return strings.Join([]string{"BrandBrokens", string(data)}, " ")
}
