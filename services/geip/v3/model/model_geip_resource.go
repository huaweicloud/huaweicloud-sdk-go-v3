package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeipResource 查询过滤标签的资源
type GeipResource struct {

	// 资源ID标识符。
	ResourceId string `json:"resource_id"`

	// 资源详情。
	ResourceDetail *interface{} `json:"resource_detail"`

	// 包含标签。
	Tags []CreateTag `json:"tags"`

	// 实例名字。
	ResourceName string `json:"resource_name"`
}

func (o GeipResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeipResource struct{}"
	}

	return strings.Join([]string{"GeipResource", string(data)}, " ")
}
