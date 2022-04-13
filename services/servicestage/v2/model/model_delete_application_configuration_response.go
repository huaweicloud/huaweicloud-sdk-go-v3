package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteApplicationConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationConfigurationResponse", string(data)}, " ")
}
