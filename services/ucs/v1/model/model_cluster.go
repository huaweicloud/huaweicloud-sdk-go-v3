package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Cluster struct {

	// API类型，固定值“Cluster”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *ClusterSpec `json:"spec,omitempty"`

	Status *ClusterStatus `json:"status,omitempty"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}
