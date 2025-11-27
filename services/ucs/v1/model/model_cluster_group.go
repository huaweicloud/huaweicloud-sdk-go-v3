package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterGroup struct {

	// API类型，固定值\"ClusterGroup\"，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *ClusterGroupSpec `json:"spec,omitempty"`

	Status *ClusterGroupStatus `json:"status,omitempty"`
}

func (o ClusterGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterGroup struct{}"
	}

	return strings.Join([]string{"ClusterGroup", string(data)}, " ")
}
