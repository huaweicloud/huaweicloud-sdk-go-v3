package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentResponse Response Object
type CreateEnvironmentResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类
	Kind *string `json:"kind,omitempty"`

	Metadata       *EnvironmentMetadata `json:"metadata,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateEnvironmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentResponse", string(data)}, " ")
}
