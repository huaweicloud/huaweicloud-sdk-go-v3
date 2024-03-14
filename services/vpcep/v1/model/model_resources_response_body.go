package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourcesResponseBody 查询配额接口响应结构体
type ResourcesResponseBody struct {

	// 资源列表
	Resources *[]Quotas `json:"resources,omitempty"`
}

func (o ResourcesResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesResponseBody struct{}"
	}

	return strings.Join([]string{"ResourcesResponseBody", string(data)}, " ")
}
