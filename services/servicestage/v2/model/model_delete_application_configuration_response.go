package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationConfigurationResponse Response Object
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
