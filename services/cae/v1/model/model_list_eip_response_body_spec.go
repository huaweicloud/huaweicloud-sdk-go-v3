package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipResponseBodySpec 出入网公网IP信息。
type ListEipResponseBodySpec struct {
	Egress *ListEipResponseBodySpecEgress `json:"egress,omitempty"`

	Ingress *ListEipResponseBodySpecIngress `json:"ingress,omitempty"`
}

func (o ListEipResponseBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipResponseBodySpec struct{}"
	}

	return strings.Join([]string{"ListEipResponseBodySpec", string(data)}, " ")
}
