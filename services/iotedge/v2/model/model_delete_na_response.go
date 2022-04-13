package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteNaResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNaResponse struct{}"
	}

	return strings.Join([]string{"DeleteNaResponse", string(data)}, " ")
}
