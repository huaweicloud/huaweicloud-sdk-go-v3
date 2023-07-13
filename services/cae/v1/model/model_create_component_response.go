package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentResponse Response Object
type CreateComponentResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Component”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
