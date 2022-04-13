package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InputResponse struct {
	// 参数名

	Name *string `json:"name,omitempty"`

	PropertyReference *PropertyReferenceResponse `json:"property_reference,omitempty"`
}

func (o InputResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputResponse struct{}"
	}

	return strings.Join([]string{"InputResponse", string(data)}, " ")
}
