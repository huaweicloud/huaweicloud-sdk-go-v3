package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowReposTagRespV3 struct {

	// tag编号
	Id int64 `json:"id"`

	// 仓库编号
	RepoId int64 `json:"repo_id"`

	// tag版本名称
	Tag string `json:"tag"`

	// 镜像id
	ImageId string `json:"image_id"`

	// 镜像manifest
	Manifest string `json:"manifest"`

	// 镜像hash值
	Digest string `json:"digest"`

	// docker协议版本，值为1或2
	Schema int64 `json:"schema"`

	// 镜像pull地址，格式为swr.cn-north-1.myhuaweicloud.com/namespace/repository:tag
	Path string `json:"path"`

	// cce集群内镜像pull路径，格式为 10.125.0.198:20202/namespace/repository:tag
	InternalPath string `json:"internal_path"`

	// 镜像大小，0 ~ 9223372036854775807
	Size int64 `json:"size"`

	// 默认值为false
	IsTrusted bool `json:"is_trusted"`

	// 镜像创建时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Created string `json:"created"`

	// 镜像更新时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Updated string `json:"updated"`

	// 帐号ID
	DomainId string `json:"domain_id"`

	// 0：manifest类型；1：manifest list类型
	TagType int64 `json:"tag_type"`
}

func (o ShowReposTagRespV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReposTagRespV3 struct{}"
	}

	return strings.Join([]string{"ShowReposTagRespV3", string(data)}, " ")
}
