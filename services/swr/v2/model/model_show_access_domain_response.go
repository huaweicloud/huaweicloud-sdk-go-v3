package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAccessDomainResponse struct {
	// true：存在；false：不存在

	Exist *bool `json:"exist,omitempty"`
	// 组织名称

	Namespace *string `json:"namespace,omitempty"`
	// 镜像仓库名称

	Repository *string `json:"repository,omitempty"`
	// 共享帐号名

	AccessDomain *string `json:"access_domain,omitempty"`
	// 权限

	Permit *string `json:"permit,omitempty"`
	// 截止时间

	Deadline *string `json:"deadline,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 创建者ID

	CreatorId *string `json:"creator_id,omitempty"`
	// 创建者名称

	CreatorName *string `json:"creator_name,omitempty"`
	// 镜像创建时间，UTC时间格式，时间为UTC标准时间

	Created *string `json:"created,omitempty"`
	// 镜像更新时间，UTC时间格式，时间为UTC标准时间

	Updated *string `json:"updated,omitempty"`
	// 是否过期，true：有效；false：过期

	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowAccessDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessDomainResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessDomainResponse", string(data)}, " ")
}
