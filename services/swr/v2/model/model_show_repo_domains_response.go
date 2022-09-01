package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowRepoDomainsResponse struct {

	// 组织
	Namespace string `json:"namespace" xml:"namespace"`

	// 镜像仓库
	Repository string `json:"repository" xml:"repository"`

	// 共享租户名
	AccessDomain string `json:"access_domain" xml:"access_domain"`

	// 权限
	Permit string `json:"permit" xml:"permit"`

	// 截止时间
	Deadline string `json:"deadline" xml:"deadline"`

	// 描述
	Description string `json:"description" xml:"description"`

	// 创建者ID
	CreatorId string `json:"creator_id" xml:"creator_id"`

	// 创建者名称
	CreatorName string `json:"creator_name" xml:"creator_name"`

	// 镜像创建时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Created string `json:"created" xml:"created"`

	// 镜像更新时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Updated string `json:"updated" xml:"updated"`

	// 是否过期：true:有效；false:过期
	Status bool `json:"status" xml:"status"`
}

func (o ShowRepoDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoDomainsResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoDomainsResponse", string(data)}, " ")
}
