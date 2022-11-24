package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteComponentConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteComponentConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComponentConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteComponentConfigurationResponse", string(data)}, " ")
}
