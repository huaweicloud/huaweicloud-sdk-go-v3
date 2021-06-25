package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAssetCategoryResponse struct {
	// 视频分类名称<br/>

	Name *string `json:"name,omitempty"`
	// 父分类ID<br/>

	ParentId *int32 `json:"parent_id,omitempty"`
	// 视频分类ID<br/>

	Id *int32 `json:"id,omitempty"`
	// 视频分类层级<br/>

	Level *int32 `json:"level,omitempty"`

	ProjectId      *string `json:"projectId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetCategoryResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetCategoryResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetCategoryResponse", string(data)}, " ")
}
