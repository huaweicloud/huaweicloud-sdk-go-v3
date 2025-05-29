package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagResource struct {

	// 证书的资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]ScsResourceTag `json:"tags,omitempty"`

	// 资源名称，默认为空字符串。
	ResourceName *string `json:"resource_name,omitempty"`

	ResourceDetail *TagResourceResourceDetail `json:"resource_detail,omitempty"`
}

func (o TagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResource struct{}"
	}

	return strings.Join([]string{"TagResource", string(data)}, " ")
}
