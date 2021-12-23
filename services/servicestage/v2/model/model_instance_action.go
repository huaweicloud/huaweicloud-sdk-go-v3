package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceAction struct {
	Action *InstanceActionType `json:"action"`

	Parameters *InstanceActionParameters `json:"parameters,omitempty"`
}

func (o InstanceAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceAction struct{}"
	}

	return strings.Join([]string{"InstanceAction", string(data)}, " ")
}
