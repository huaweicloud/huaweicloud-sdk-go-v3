package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeApplicationConfigurationResponse struct {
	// 应用ID。

	ApplicationId *string `json:"application_id,omitempty"`
	// 环境ID。

	EnvironmentId *string `json:"environment_id,omitempty"`

	Configuration  *ApplicationListConfigConfiguration `json:"configuration,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ChangeApplicationConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeApplicationConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ChangeApplicationConfigurationResponse", string(data)}, " ")
}
