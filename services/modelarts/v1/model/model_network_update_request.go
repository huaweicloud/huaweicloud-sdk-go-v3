package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkUpdateRequest 更新网络资源的请求体。
type NetworkUpdateRequest struct {
	Metadata *NetworkMetadataUpdate `json:"metadata,omitempty"`

	Spec *NetworkSpecUpdate `json:"spec,omitempty"`
}

func (o NetworkUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkUpdateRequest struct{}"
	}

	return strings.Join([]string{"NetworkUpdateRequest", string(data)}, " ")
}
