package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowReposResp struct {

	// 仓库名称
	Name string `json:"name" xml:"name"`

	// 仓库类型（计划改造，每个镜像会有多个lable标示）
	Category string `json:"category" xml:"category"`

	// 仓库描述信息
	Description string `json:"description" xml:"description"`

	// 仓库大小
	Size int64 `json:"size" xml:"size"`

	// 仓库是否为公共仓库，值为true或false
	IsPublic bool `json:"is_public" xml:"is_public"`

	// 仓库中镜像个数，0 ~ 9223372036854775807
	NumImages int64 `json:"num_images" xml:"num_images"`

	// 仓库下载次数
	NumDownload int64 `json:"num_download" xml:"num_download"`

	// 仓库创建时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	CreatedAt string `json:"created_at" xml:"created_at"`

	// 仓库更新时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	UpdatedAt string `json:"updated_at" xml:"updated_at"`

	// 仓库logo地址（暂时未用）
	Logo string `json:"logo" xml:"logo"`

	// 仓库logo图片的URL，URL格式。（暂时未用）
	Url string `json:"url" xml:"url"`

	// 镜像pull路径，格式为 swr.cn-north-1.myhuaweicloud.com/namespace/repository
	Path string `json:"path" xml:"path"`

	// 镜像pull路径，格式为 10.125.0.198:20202/namespace/repository
	InternalPath string `json:"internal_path" xml:"internal_path"`

	// 租户名
	DomainName string `json:"domain_name" xml:"domain_name"`

	// 租户的组织名称
	Namespace string `json:"namespace" xml:"namespace"`

	// 镜像版本列表
	Tags []string `json:"tags" xml:"tags"`

	// 查询他人共享镜像：共享是否过期 查询我共享的镜像：默认为false,无意义
	Status bool `json:"status" xml:"status"`

	// 总记录条数
	TotalRange int64 `json:"total_range" xml:"total_range"`
}

func (o ShowReposResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReposResp struct{}"
	}

	return strings.Join([]string{"ShowReposResp", string(data)}, " ")
}
