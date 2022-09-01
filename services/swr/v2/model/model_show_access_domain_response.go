package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAccessDomainResponse struct {

	// true：存在；false：不存在
	Exist *bool `json:"exist,omitempty" xml:"exist"`

	// 组织名称
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 镜像仓库名称
	Repository *string `json:"repository,omitempty" xml:"repository"`

	// 共享帐号名
	AccessDomain *string `json:"access_domain,omitempty" xml:"access_domain"`

	// 权限
	Permit *string `json:"permit,omitempty" xml:"permit"`

	// 截止时间
	Deadline *string `json:"deadline,omitempty" xml:"deadline"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 镜像创建时间，UTC时间格式，时间为UTC标准时间
	Created *string `json:"created,omitempty" xml:"created"`

	// 镜像更新时间，UTC时间格式，时间为UTC标准时间
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 是否过期，true：有效；false：过期
	Status         *bool `json:"status,omitempty" xml:"status"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowAccessDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessDomainResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessDomainResponse", string(data)}, " ")
}
