package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTempResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	// message
	Message        *string `json:"message,omitempty" xml:"message"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTempResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTempResponse struct{}"
	}

	return strings.Join([]string{"DeleteTempResponse", string(data)}, " ")
}
