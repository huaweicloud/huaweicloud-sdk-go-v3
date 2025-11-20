package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedDn struct {

	// DN实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// DN实例名称。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o RelatedDn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedDn struct{}"
	}

	return strings.Join([]string{"RelatedDn", string(data)}, " ")
}
