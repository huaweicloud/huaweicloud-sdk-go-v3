package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsResource struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称，没有默认为空字符串
	ResourceName *string `json:"resource_name,omitempty"`

	ResourceDetail *ResourceDetail `json:"resource_detail,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o TmsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResource struct{}"
	}

	return strings.Join([]string{"TmsResource", string(data)}, " ")
}
