package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnvironmentResponse Response Object
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
