package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFactorParam struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Type *string `json:"type,omitempty"`

	Name *string `json:"name,omitempty"`

	ParentNodeIds *[]string `json:"parent_node_ids,omitempty"`

	AssetId *string `json:"asset_id,omitempty"`

	CreatorNum *string `json:"creator_num,omitempty"`

	MindmapId *string `json:"mindmap_id,omitempty"`

	TestpointId *string `json:"testpoint_id,omitempty"`

	MindmapNodeId *string `json:"mindmap_node_id,omitempty"`
}

func (o ListFactorParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactorParam struct{}"
	}

	return strings.Join([]string{"ListFactorParam", string(data)}, " ")
}
