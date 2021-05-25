package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceCreateTagRequest struct {
	// 镜像id

	ImageId string `json:"image_id"`
	// 新增的tag。字符串中不能包含“=”和“.”。

	Tag string `json:"tag"`
}

func (o GlanceCreateTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceCreateTagRequest struct{}"
	}

	return strings.Join([]string{"GlanceCreateTagRequest", string(data)}, " ")
}
