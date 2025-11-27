package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Pod struct {
	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *PodSpec `json:"spec,omitempty"`

	Status *PodStatus `json:"status,omitempty"`
}

func (o Pod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pod struct{}"
	}

	return strings.Join([]string{"Pod", string(data)}, " ")
}
