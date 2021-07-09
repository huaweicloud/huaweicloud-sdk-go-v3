package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceParamRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`
	// 语种，默认中文。中文:zh-cn;英文:en-us

	XLanguage *string `json:"X-Language,omitempty"`

	Body *UpdateParametersReq `json:"body,omitempty"`
}

func (o UpdateInstanceParamRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceParamRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceParamRequest", string(data)}, " ")
}
