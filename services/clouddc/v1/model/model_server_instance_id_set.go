package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerInstanceIdSet 实例ID集合
type ServerInstanceIdSet struct {

	// instance id set
	InstanceIdSet *[]string `json:"instance_id_set,omitempty"`
}

func (o ServerInstanceIdSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerInstanceIdSet struct{}"
	}

	return strings.Join([]string{"ServerInstanceIdSet", string(data)}, " ")
}
