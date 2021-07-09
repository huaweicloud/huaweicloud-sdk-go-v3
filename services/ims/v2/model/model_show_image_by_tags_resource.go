package model

import (
	"encoding/json"

	"strings"
)

// resource字段数据结构说明
type ShowImageByTagsResource struct {
	// 镜像ID

	ResourceId string `json:"resource_id"`

	ResourceDetail *QueryImageByTagsResourceDetail `json:"resource_detail"`
	// 镜像的标签列表

	Tags []TagKeyValue `json:"tags"`
	// 镜像名称

	ResourceName string `json:"resource_name"`
}

func (o ShowImageByTagsResource) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowImageByTagsResource struct{}"
	}

	return strings.Join([]string{"ShowImageByTagsResource", string(data)}, " ")
}
