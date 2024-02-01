package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmsResource struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 标签列表，没有标签默认为空数组。
	Tags []Tag `json:"tags"`

	// 资源名称，资源没有名称时默认为空字符串。
	ResourceName string `json:"resource_name"`
}

func (o TmsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResource struct{}"
	}

	return strings.Join([]string{"TmsResource", string(data)}, " ")
}
