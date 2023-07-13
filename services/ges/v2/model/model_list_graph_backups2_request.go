package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGraphBackups2Request Request Object
type ListGraphBackups2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 本次请求的起始位置，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页资源数量的最大值，默认为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListGraphBackups2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphBackups2Request struct{}"
	}

	return strings.Join([]string{"ListGraphBackups2Request", string(data)}, " ")
}
