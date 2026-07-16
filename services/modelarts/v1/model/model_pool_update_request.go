package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolUpdateRequest 更新网络资源的请求体。
type PoolUpdateRequest struct {
	Metadata *PoolMetadataUpdate `json:"metadata,omitempty"`

	Spec *PoolSpecUpdate `json:"spec,omitempty"`
}

func (o PoolUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolUpdateRequest struct{}"
	}

	return strings.Join([]string{"PoolUpdateRequest", string(data)}, " ")
}
