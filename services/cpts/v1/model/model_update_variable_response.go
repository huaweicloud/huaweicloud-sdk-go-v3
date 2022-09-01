package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateVariableResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	Json *CreateVariableResultJson `json:"json,omitempty" xml:"json"`

	// message
	Message        *string `json:"message,omitempty" xml:"message"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVariableResponse struct{}"
	}

	return strings.Join([]string{"UpdateVariableResponse", string(data)}, " ")
}
