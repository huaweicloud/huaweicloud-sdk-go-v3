package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationResponse Response Object
type CreateApplicationResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ApplicationKindObj `json:"kind,omitempty"`

	Metadata       *ApplicationMetadata `json:"metadata,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationResponse", string(data)}, " ")
}
