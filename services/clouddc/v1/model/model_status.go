package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Status 硬件状态总览
type Status struct {
	State *State `json:"state,omitempty"`

	Health *Health `json:"health,omitempty"`
}

func (o Status) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Status struct{}"
	}

	return strings.Join([]string{"Status", string(data)}, " ")
}
