package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcEgressResponse Response Object
type ListVpcEgressResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *VpcEgressKindObj `json:"kind,omitempty"`

	Spec           *VpcEgressResponseBodySpec `json:"spec,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListVpcEgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcEgressResponse struct{}"
	}

	return strings.Join([]string{"ListVpcEgressResponse", string(data)}, " ")
}
