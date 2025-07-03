package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepoTagResponse Response Object
type ShowRepoTagResponse struct {

	// tag编号
	Id *int64 `json:"id,omitempty"`

	// 仓库编号
	RepoId *int64 `json:"repo_id,omitempty"`

	// tag版本名称
	Tag *string `json:"tag,omitempty"`

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 镜像manifest
	Manifest *string `json:"manifest,omitempty"`

	// 镜像hash值
	Digest *string `json:"digest,omitempty"`

	// docker协议版本，值为1或2
	Schema *int64 `json:"schema,omitempty"`

	// 镜像pull地址，格式为swr.cn-north-1.myhuaweicloud.com/namespace/repository:tag
	Path *string `json:"path,omitempty"`

	// cce集群内镜像pull路径，格式为 10.125.0.198:20202/namespace/repository:tag
	InternalPath *string `json:"internal_path,omitempty"`

	// 镜像大小，0 ~ 9223372036854775807
	Size *int64 `json:"size,omitempty"`

	// 默认值为false
	IsTrusted *bool `json:"is_trusted,omitempty"`

	// 镜像创建时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Created *string `json:"created,omitempty"`

	// 镜像更新时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Updated *string `json:"updated,omitempty"`

	// 帐号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 0：manifest类型；1：manifest list类型
	TagType        *int64 `json:"tag_type,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRepoTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoTagResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoTagResponse", string(data)}, " ")
}
