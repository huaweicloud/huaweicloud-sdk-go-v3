package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterGroupStatus struct {
	Conditions *[]ClusterGroupCondition `json:"conditions,omitempty"`
}

func (o ClusterGroupStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterGroupStatus struct{}"
	}

	return strings.Join([]string{"ClusterGroupStatus", string(data)}, " ")
}
