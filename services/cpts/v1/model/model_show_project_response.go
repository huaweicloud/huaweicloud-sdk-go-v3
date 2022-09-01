package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProjectResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	// message
	Message *string `json:"message,omitempty" xml:"message"`

	Project        *Project `json:"project,omitempty" xml:"project"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectResponse", string(data)}, " ")
}
