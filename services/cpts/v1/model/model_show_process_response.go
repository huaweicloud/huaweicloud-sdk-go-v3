package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProcessResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	// message
	Message *string `json:"message,omitempty" xml:"message"`

	Json *UploadProcessJson `json:"json,omitempty" xml:"json"`

	// extend
	Extend         *string `json:"extend,omitempty" xml:"extend"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProcessResponse struct{}"
	}

	return strings.Join([]string{"ShowProcessResponse", string(data)}, " ")
}
