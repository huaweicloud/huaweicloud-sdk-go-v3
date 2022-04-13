package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteServerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerPasswordResponse", string(data)}, " ")
}
