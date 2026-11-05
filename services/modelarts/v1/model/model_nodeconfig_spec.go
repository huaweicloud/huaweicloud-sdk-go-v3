package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeconfigSpec struct {
	Configs *NodeconfigSpecConfigs `json:"configs"`
}

func (o NodeconfigSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigSpec struct{}"
	}

	return strings.Join([]string{"NodeconfigSpec", string(data)}, " ")
}
