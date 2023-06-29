package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LeftRightPosition
type LeftRightPosition struct {
	Left *Position `json:"left"`

	Right *Position `json:"right"`
}

func (o LeftRightPosition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeftRightPosition struct{}"
	}

	return strings.Join([]string{"LeftRightPosition", string(data)}, " ")
}
