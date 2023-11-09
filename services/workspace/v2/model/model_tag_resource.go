package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagResource struct {

	// 计算机信息
	ResourceDetail *string `json:"resource_detail,omitempty"`

	// 计算机id
	ResourceId *string `json:"resource_id,omitempty"`

	// 计算机名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 标签对象
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o TagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResource struct{}"
	}

	return strings.Join([]string{"TagResource", string(data)}, " ")
}
