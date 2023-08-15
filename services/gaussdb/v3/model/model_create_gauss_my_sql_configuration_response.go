package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGaussMySqlConfigurationResponse Response Object
type CreateGaussMySqlConfigurationResponse struct {
	Configurations *ConfigurationSummary2 `json:"configurations,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateGaussMySqlConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlConfigurationResponse", string(data)}, " ")
}
