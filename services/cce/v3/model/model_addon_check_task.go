package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddonCheckTask struct {
	Metadata *CheckTaskMetadata `json:"metadata"`

	Spec *CheckTaskSpec `json:"spec,omitempty"`

	Status *CheckTaskStatus `json:"status"`
}

func (o AddonCheckTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonCheckTask struct{}"
	}

	return strings.Join([]string{"AddonCheckTask", string(data)}, " ")
}
