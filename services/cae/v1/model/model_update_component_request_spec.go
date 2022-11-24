package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateComponentRequestSpec struct {
	Source *Source `json:"source,omitempty"`

	Build *Build `json:"build,omitempty"`

	ResourceLimit *ResourceLimit `json:"resource_limit"`

	LogStrategy *LogStrategy `json:"log_strategy,omitempty"`
}

func (o UpdateComponentRequestSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequestSpec struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequestSpec", string(data)}, " ")
}
