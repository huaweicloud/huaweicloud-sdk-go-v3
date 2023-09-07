package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentResponse Response Object
type CreateComponentResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentKindObj `json:"kind,omitempty"`

	Metadata *MetadataResponse `json:"metadata,omitempty"`

	Spec           *CreateComponentSpec `json:"spec,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentResponse struct{}"
	}

	return strings.Join([]string{"CreateComponentResponse", string(data)}, " ")
}
