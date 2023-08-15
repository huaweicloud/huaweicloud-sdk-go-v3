package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryResponse Response Object
type ShowRepositoryResponse struct {

	// 仓库编号
	Id *int64 `json:"id,omitempty"`

	// 组织编号
	NsId *int64 `json:"ns_id,omitempty"`

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 仓库类型（计划改造，每个镜像会有多个lable标示）
	Category *string `json:"category,omitempty"`

	// 仓库描述信息
	Description *string `json:"description,omitempty"`

	// 仓库创建者id
	CreatorId *string `json:"creator_id,omitempty"`

	// 仓库创建者
	CreatorName *string `json:"creator_name,omitempty"`

	// 仓库大小
	Size *int64 `json:"size,omitempty"`

	// 仓库是否为公共仓库，值为true或false
	IsPublic *bool `json:"is_public,omitempty"`

	// 仓库中镜像个数，0 ~ 9223372036854775807
	NumImages *int64 `json:"num_images,omitempty"`

	// 仓库下载次数
	NumDownload *int64 `json:"num_download,omitempty"`

	// 仓库logo图片的URL，URL格式。（暂时未用）
	Url *string `json:"url,omitempty"`

	// 镜像pull路径，格式为 swr.cn-north-1.myhuaweicloud.com/namespace/repository
	Path *string `json:"path,omitempty"`

	// 镜像pull路径，格式为 10.125.0.198:20202/namespace/repository
	InternalPath *string `json:"internal_path,omitempty"`

	// 仓库创建时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Created *string `json:"created,omitempty"`

	// 仓库更新时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Updated *string `json:"updated,omitempty"`

	// 帐号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 镜像排序优先级
	Priority       *int32 `json:"priority,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryResponse", string(data)}, " ")
}
