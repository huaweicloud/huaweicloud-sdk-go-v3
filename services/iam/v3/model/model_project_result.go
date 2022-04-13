package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ProjectResult struct {
	// false.

	IsDomain bool `json:"is_domain"`
	// 项目描述信息。

	Description string `json:"description"`

	Links *Links `json:"links"`
	// 项目是否可用。

	Enabled bool `json:"enabled"`
	// 项目ID。

	Id string `json:"id"`
	// 如果查询自己创建的项目，则此处返回所属区域的项目ID。  如果查询的是系统内置项目，如cn-north-4，则此处返回账号ID。

	ParentId string `json:"parent_id"`
	// 项目所属账号ID。

	DomainId string `json:"domain_id"`
	// 项目名称。

	Name string `json:"name"`
}

func (o ProjectResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectResult struct{}"
	}

	return strings.Join([]string{"ProjectResult", string(data)}, " ")
}
