package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCasePageParam struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Deleted *string `json:"deleted,omitempty"`

	IdCollection *[]string `json:"id_collection,omitempty"`

	MindmapId *string `json:"mindmap_id,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NodeIdCollection *[]string `json:"node_id_collection,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	IsArchive *bool `json:"is_archive,omitempty"`

	CaseName *string `json:"case_name,omitempty"`

	HasSubMindmap *bool `json:"has_sub_mindmap,omitempty"`

	SubMindmapId *[]string `json:"sub_mindmap_id,omitempty"`
}

func (o TestCasePageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCasePageParam struct{}"
	}

	return strings.Join([]string{"TestCasePageParam", string(data)}, " ")
}
