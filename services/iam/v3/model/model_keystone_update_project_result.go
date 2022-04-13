package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KeystoneUpdateProjectResult struct {
	// false.

	IsDomain bool `json:"is_domain"`
	// 项目描述信息。

	Description string `json:"description"`
	// 项目的其他信息。

	Extra *interface{} `json:"extra,omitempty"`

	Links *LinksSelf `json:"links"`
	// 项目是否可用。

	Enabled bool `json:"enabled"`
	// 项目ID。

	Id string `json:"id"`
	// 区域对应的项目ID，例如区域“华北-北京一”区域对应的项目ID为：04dd42abe48026ad2fa3c01ad7fa.....。

	ParentId string `json:"parent_id"`
	// 项目所属账号ID。

	DomainId string `json:"domain_id"`
	// 项目名称。

	Name string `json:"name"`
}

func (o KeystoneUpdateProjectResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProjectResult struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProjectResult", string(data)}, " ")
}
