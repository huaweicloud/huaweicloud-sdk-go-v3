package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterEndpointsResponse Response Object
type ShowClusterEndpointsResponse struct {
	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *OpenApiSpec `json:"spec,omitempty"`

	Status         *OpenApiResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowClusterEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterEndpointsResponse", string(data)}, " ")
}
