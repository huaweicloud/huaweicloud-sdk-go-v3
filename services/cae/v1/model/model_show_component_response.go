package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComponentResponse Response Object
type ShowComponentResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Component”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
