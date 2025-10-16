package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceIdsStatus struct {
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	InstanceIds *[]string `json:"instance_ids,omitempty"`
}

func (o InstanceIdsStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceIdsStatus struct{}"
	}

	return strings.Join([]string{"InstanceIdsStatus", string(data)}, " ")
}
