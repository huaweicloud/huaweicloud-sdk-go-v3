package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type State struct {
	Phase string `json:"phase"`

	Reason *string `json:"reason,omitempty"`
}

func (o State) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "State struct{}"
	}

	return strings.Join([]string{"State", string(data)}, " ")
}
