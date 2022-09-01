package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源结构体。
type TagResource struct {

	// 资源ID。
	ResourceId string `json:"resource_id" xml:"resource_id"`

	ResourceDetail *ResourceDetail `json:"resource_detail" xml:"resource_detail"`

	// 标签列表。
	Tags []ResourceTag `json:"tags" xml:"tags"`

	// 资源名称。
	ResourceName string `json:"resource_name" xml:"resource_name"`
}

func (o TagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResource struct{}"
	}

	return strings.Join([]string{"TagResource", string(data)}, " ")
}
