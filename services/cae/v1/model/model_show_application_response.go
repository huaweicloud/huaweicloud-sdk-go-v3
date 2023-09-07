package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplicationResponse Response Object
type ShowApplicationResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ApplicationKindObj `json:"kind,omitempty"`

	Metadata       *ApplicationMetadata `json:"metadata,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationResponse struct{}"
	}

	return strings.Join([]string{"ShowApplicationResponse", string(data)}, " ")
}
