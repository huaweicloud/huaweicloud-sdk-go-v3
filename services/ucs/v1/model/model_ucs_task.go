package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsTask struct {

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *UcsTaskSpec `json:"spec,omitempty"`

	Status *UcsTaskStatus `json:"status,omitempty"`
}

func (o UcsTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsTask struct{}"
	}

	return strings.Join([]string{"UcsTask", string(data)}, " ")
}
