package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMetadataReqGesMetadataLabels struct {

	// label名称
	Name *string `json:"name,omitempty"`

	// label属性map
	Properties *[]map[string]string `json:"properties,omitempty"`
}

func (o CreateMetadataReqGesMetadataLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataReqGesMetadataLabels struct{}"
	}

	return strings.Join([]string{"CreateMetadataReqGesMetadataLabels", string(data)}, " ")
}
