package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcEgressRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *VpcEgressKindObj `json:"kind,omitempty"`

	Spec *VpcEgressRequestBodySpec `json:"spec,omitempty"`
}

func (o VpcEgressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcEgressRequestBody struct{}"
	}

	return strings.Join([]string{"VpcEgressRequestBody", string(data)}, " ")
}
