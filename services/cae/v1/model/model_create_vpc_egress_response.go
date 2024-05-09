package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcEgressResponse Response Object
type CreateVpcEgressResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *VpcEgressKindObj `json:"kind,omitempty"`

	Spec           *VpcEgressResponseBodySpec `json:"spec,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreateVpcEgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcEgressResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcEgressResponse", string(data)}, " ")
}
