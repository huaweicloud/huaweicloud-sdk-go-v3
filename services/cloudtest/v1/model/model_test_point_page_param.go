package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestPointPageParam struct {
	Deleted *string `json:"deleted,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	MindmapId *string `json:"mindmap_id,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o TestPointPageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPointPageParam struct{}"
	}

	return strings.Join([]string{"TestPointPageParam", string(data)}, " ")
}
