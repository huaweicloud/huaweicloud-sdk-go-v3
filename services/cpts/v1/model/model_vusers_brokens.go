package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VusersBrokens struct {
	// vusers

	Vusers *[]float64 `json:"vusers,omitempty"`
}

func (o VusersBrokens) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VusersBrokens struct{}"
	}

	return strings.Join([]string{"VusersBrokens", string(data)}, " ")
}
