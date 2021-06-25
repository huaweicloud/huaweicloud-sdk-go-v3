package model

import (
	"encoding/json"

	"strings"
)

type UpdateCategoryReq struct {
	// 视频分类名称

	Name string `json:"name"`
	// 视频分类ID

	Id int32 `json:"id"`
}

func (o UpdateCategoryReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCategoryReq struct{}"
	}

	return strings.Join([]string{"UpdateCategoryReq", string(data)}, " ")
}
