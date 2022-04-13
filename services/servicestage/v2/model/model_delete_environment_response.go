package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEnvironmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentResponse", string(data)}, " ")
}
