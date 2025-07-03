package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowReposRespV3 struct {

	// 仓库ID
	Id int64 `json:"id"`

	// 仓库名称
	Name string `json:"name"`

	// 仓库类型（计划改造，每个镜像会有多个lable标示）
	Category string `json:"category"`

	// 仓库描述信息
	Description string `json:"description"`

	// 仓库大小
	Size int64 `json:"size"`

	// 仓库是否为公共仓库，值为true或false
	IsPublic bool `json:"is_public"`

	// 仓库中镜像个数，0 ~ 9223372036854775807
	NumImages int64 `json:"num_images"`

	// 仓库下载次数
	NumDownload int64 `json:"num_download"`

	// 仓库创建时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	CreatedAt string `json:"created_at"`

	// 仓库更新时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	UpdatedAt string `json:"updated_at"`

	// 仓库所属租户
	DomainName string `json:"domain_name"`

	// 租户的组织名称
	Namespace string `json:"namespace"`

	// 查询他人共享镜像：共享是否过期 查询我共享的镜像：默认为false,无意义
	Status *bool `json:"status,omitempty"`
}

func (o ShowReposRespV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReposRespV3 struct{}"
	}

	return strings.Join([]string{"ShowReposRespV3", string(data)}, " ")
}
