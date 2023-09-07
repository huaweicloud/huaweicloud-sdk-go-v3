package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEipRequestBodySpec 修改出入网公网IP带宽信息。
type UpdateEipRequestBodySpec struct {
	Egress *UpdateEipRequestBodySpecEgress `json:"egress,omitempty"`

	Ingress *UpdateEipRequestBodySpecIngress `json:"ingress,omitempty"`
}

func (o UpdateEipRequestBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEipRequestBodySpec struct{}"
	}

	return strings.Join([]string{"UpdateEipRequestBodySpec", string(data)}, " ")
}
