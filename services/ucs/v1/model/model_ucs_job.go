package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsJob struct {

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *UcsJobSpec `json:"spec,omitempty"`

	Status *UcsJobStatus `json:"status,omitempty"`
}

func (o UcsJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsJob struct{}"
	}

	return strings.Join([]string{"UcsJob", string(data)}, " ")
}
