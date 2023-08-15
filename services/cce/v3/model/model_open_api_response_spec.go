package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenApiResponseSpec struct {
	Spec *OpenApiResponseSpecSpec `json:"spec,omitempty"`
}

func (o OpenApiResponseSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApiResponseSpec struct{}"
	}

	return strings.Join([]string{"OpenApiResponseSpec", string(data)}, " ")
}
