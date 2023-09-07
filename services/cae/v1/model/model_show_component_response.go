package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComponentResponse Response Object
type ShowComponentResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentKindObj `json:"kind,omitempty"`

	Metadata *MetadataResponse `json:"metadata,omitempty"`

	Spec           *ComponentSpec `json:"spec,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentResponse struct{}"
	}

	return strings.Join([]string{"ShowComponentResponse", string(data)}, " ")
}
