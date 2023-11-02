package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdatableVersionRequest Request Object
type ListUpdatableVersionRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 条目数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUpdatableVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdatableVersionRequest struct{}"
	}

	return strings.Join([]string{"ListUpdatableVersionRequest", string(data)}, " ")
}
