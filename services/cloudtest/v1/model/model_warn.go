package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Warn struct {
	WarnCode *string `json:"warn_code,omitempty"`

	WarnMsg *string `json:"warn_msg,omitempty"`
}

func (o Warn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Warn struct{}"
	}

	return strings.Join([]string{"Warn", string(data)}, " ")
}
