package model

import (
	"encoding/json"

	"strings"
)

type CreateCategoryReq struct {
	// 视频分类名称

	Name string `json:"name"`
	// 父分类ID

	ParentId *int32 `json:"parent_id,omitempty"`
}

func (o CreateCategoryReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCategoryReq struct{}"
	}

	return strings.Join([]string{"CreateCategoryReq", string(data)}, " ")
}
