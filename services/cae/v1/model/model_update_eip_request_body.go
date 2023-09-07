package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEipRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *EipKindObj `json:"kind,omitempty"`

	Spec *UpdateEipRequestBodySpec `json:"spec,omitempty"`
}

func (o UpdateEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEipRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEipRequestBody", string(data)}, " ")
}
