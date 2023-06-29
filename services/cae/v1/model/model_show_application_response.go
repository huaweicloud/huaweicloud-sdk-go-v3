package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplicationResponse Response Object
type ShowApplicationResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *string `json:"kind,omitempty"`

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
