package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddonInstance struct {

	// API类型，必须是\"Addon\"
	Kind string `json:"kind"`

	// API版本，必须是\"v3\"
	ApiVersion string `json:"apiVersion"`

	Metadata *ObjectMeta `json:"metadata"`

	Spec *AddonInstanceSpec `json:"spec,omitempty"`

	Status *AddonInstanceStatus `json:"status,omitempty"`
}

func (o AddonInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonInstance struct{}"
	}

	return strings.Join([]string{"AddonInstance", string(data)}, " ")
}
