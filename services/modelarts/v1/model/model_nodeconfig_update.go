package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeconfigUpdate struct {
	Spec *NodeconfigUpdateSpec `json:"spec"`
}

func (o NodeconfigUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigUpdate struct{}"
	}

	return strings.Join([]string{"NodeconfigUpdate", string(data)}, " ")
}
