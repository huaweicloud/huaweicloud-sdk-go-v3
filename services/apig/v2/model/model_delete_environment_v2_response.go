package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEnvironmentV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnvironmentV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentV2Response struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentV2Response", string(data)}, " ")
}
