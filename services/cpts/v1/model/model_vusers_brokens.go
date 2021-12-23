package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VusersBrokens struct {
	// vusers

	Vusers *[]int32 `json:"vusers,omitempty"`
}

func (o VusersBrokens) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VusersBrokens struct{}"
	}

	return strings.Join([]string{"VusersBrokens", string(data)}, " ")
}
