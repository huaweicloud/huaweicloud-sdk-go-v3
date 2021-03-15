package model

import (
	"encoding/json"

	"strings"
)

// 镜像详情
type QueryImageByTagsResourceDetail struct {
	// 镜像状态
	Status string `json:"status"`
}

func (o QueryImageByTagsResourceDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryImageByTagsResourceDetail struct{}"
	}

	return strings.Join([]string{"QueryImageByTagsResourceDetail", string(data)}, " ")
}
