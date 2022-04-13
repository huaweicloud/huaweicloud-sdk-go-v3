package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateIpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpResponse struct{}"
	}

	return strings.Join([]string{"CreateIpResponse", string(data)}, " ")
}
