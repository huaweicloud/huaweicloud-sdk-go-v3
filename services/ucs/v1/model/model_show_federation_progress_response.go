package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFederationProgressResponse Response Object
type ShowFederationProgressResponse struct {

	// API类型。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *JobSpec `json:"spec,omitempty"`

	Status         *JobStatus `json:"status,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowFederationProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFederationProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowFederationProgressResponse", string(data)}, " ")
}
