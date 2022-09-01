package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowClusterEndpointsResponse struct {
	Metadata *Metadata `json:"metadata,omitempty" xml:"metadata"`

	Spec *OpenApiResponseSpec `json:"spec,omitempty" xml:"spec"`

	Status         *MasterEipResponseStatus `json:"status,omitempty" xml:"status"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowClusterEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterEndpointsResponse", string(data)}, " ")
}
