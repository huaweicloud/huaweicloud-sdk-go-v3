package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClusterRequestBody struct {

	// API类型，固定值“Cluster”，该值不可修改。
	Kind string `json:"kind"`

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"apiVersion"`

	Metadata *UpdateObjectMeta `json:"metadata,omitempty"`

	Spec *UpdateClusterSpec `json:"spec,omitempty"`
}

func (o UpdateClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateClusterRequestBody", string(data)}, " ")
}
