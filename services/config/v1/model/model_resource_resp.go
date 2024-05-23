package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceResp struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源详情。 资源对象，用于扩展。默认为空
	ResourceDetail *interface{} `json:"resource_detail"`

	// 标签列表，没有标签默认为空数组
	Tags []ResourceTag `json:"tags"`

	// 资源名称，资源没有名称时默认为空字符串，eip返回ip地址。
	ResourceName string `json:"resource_name"`
}

func (o ResourceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceResp struct{}"
	}

	return strings.Join([]string{"ResourceResp", string(data)}, " ")
}
