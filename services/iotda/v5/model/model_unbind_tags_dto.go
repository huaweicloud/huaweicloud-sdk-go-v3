package model

import (
	"encoding/json"

	"strings"
)

// 创建标签请求结构体。
type UnbindTagsDto struct {
	// 要解绑标签的资源类型。 - device：设备。
	ResourceType string `json:"resource_type"`
	// 要解绑标签的资源id。例如，资源类型为device，那么对应的资源id就是device_id。
	ResourceId string `json:"resource_id"`
	// 指定资源要解绑的标签键列表，标签键列表中各项之间不允许重复，不能填写不存在的标签键值。
	TagKeys []string `json:"tag_keys"`
}

func (o UnbindTagsDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnbindTagsDto struct{}"
	}

	return strings.Join([]string{"UnbindTagsDto", string(data)}, " ")
}
