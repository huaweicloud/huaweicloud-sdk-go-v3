package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionResources struct {
	// 资源ID。

	ResourceId *string `json:"resource_id,omitempty"`

	ResourceDetail *KeyDetails `json:"resource_detail,omitempty"`
	// 资源名称，默认为空字符串。

	ResourceName *string `json:"resource_name,omitempty"`
	// 标签列表，没有标签，数组默认为空。

	Tags *[]TagItem `json:"tags,omitempty"`
}

func (o ActionResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionResources struct{}"
	}

	return strings.Join([]string{"ActionResources", string(data)}, " ")
}
