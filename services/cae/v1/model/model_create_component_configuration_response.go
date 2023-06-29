package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentConfigurationResponse Response Object
type CreateComponentConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateComponentConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateComponentConfigurationResponse", string(data)}, " ")
}
