package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyCicdConfigurationResponse Response Object
type ModifyCicdConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyCicdConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyCicdConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ModifyCicdConfigurationResponse", string(data)}, " ")
}
