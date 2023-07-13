package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationResponse Response Object
type CreateApplicationResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Application”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
