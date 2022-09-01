package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// resource字段数据结构说明
type ShowImageByTagsResource struct {

	// 镜像ID
	ResourceId string `json:"resource_id" xml:"resource_id"`

	ResourceDetail *QueryImageByTagsResourceDetail `json:"resource_detail" xml:"resource_detail"`

	// 镜像的标签列表
	Tags []TagKeyValue `json:"tags" xml:"tags"`

	// 镜像名称
	ResourceName string `json:"resource_name" xml:"resource_name"`
}

func (o ShowImageByTagsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageByTagsResource struct{}"
	}

	return strings.Join([]string{"ShowImageByTagsResource", string(data)}, " ")
}
