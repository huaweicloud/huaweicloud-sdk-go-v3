package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListGraphBackupsRequest struct {
	// 图ID。

	GraphId string `json:"graph_id"`
	// 每页资源数量的最大值，默认为10。

	Limit *int32 `json:"limit,omitempty"`
	// 本次请求的起始位置，默认为0。

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListGraphBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListGraphBackupsRequest", string(data)}, " ")
}
