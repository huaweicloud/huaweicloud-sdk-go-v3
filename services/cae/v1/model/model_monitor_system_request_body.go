package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonitorSystemRequestBody struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *MonitorSystemKindObj `json:"kind,omitempty"`

	Spec *MonitorSystemRequestBodySpec `json:"spec,omitempty"`
}

func (o MonitorSystemRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorSystemRequestBody struct{}"
	}

	return strings.Join([]string{"MonitorSystemRequestBody", string(data)}, " ")
}
