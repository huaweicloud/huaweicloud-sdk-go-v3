package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Cluster struct {

	// API类型，固定值“Cluster”或“cluster”，该值不可修改。
	Kind string `json:"kind" xml:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion string `json:"apiVersion" xml:"apiVersion"`

	Metadata *ClusterMetadata `json:"metadata" xml:"metadata"`

	Spec *ClusterSpec `json:"spec" xml:"spec"`

	Status *ClusterStatus `json:"status,omitempty" xml:"status"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}
