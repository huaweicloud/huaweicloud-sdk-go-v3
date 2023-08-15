package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTempResponse Response Object
type DeleteTempResponse struct {

	// code
	Code *string `json:"code,omitempty"`

	// message
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTempResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTempResponse struct{}"
	}

	return strings.Join([]string{"DeleteTempResponse", string(data)}, " ")
}
