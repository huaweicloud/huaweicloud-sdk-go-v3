package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowReposTagResp struct {

	// tag编号
	Id int64 `json:"id" xml:"id"`

	// 仓库编号
	RepoId int64 `json:"repo_id" xml:"repo_id"`

	// Tag版本名称
	Tag string `json:"Tag" xml:"Tag"`

	// 镜像id
	ImageId string `json:"image_id" xml:"image_id"`

	// 镜像manifest
	Manifest string `json:"manifest" xml:"manifest"`

	// 镜像hash值
	Digest string `json:"digest" xml:"digest"`

	// docker协议版本，值为1或2
	Schema int64 `json:"schema" xml:"schema"`

	// 镜像pull地址，格式为swr.cn-north-1.myhuaweicloud.com/namespace/repository:tag
	Path string `json:"path" xml:"path"`

	// cce集群内镜像pull路径，格式为 10.125.0.198:20202/namespace/repository:tag
	InternalPath string `json:"internal_path" xml:"internal_path"`

	// 镜像大小，0 ~ 9223372036854775807
	Size int64 `json:"size" xml:"size"`

	// 默认值为false
	IsTrusted bool `json:"is_trusted" xml:"is_trusted"`

	// 镜像创建时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Created string `json:"created" xml:"created"`

	// 镜像更新时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Updated string `json:"updated" xml:"updated"`

	// 镜像删除时间，UTC时间格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	Deleted string `json:"deleted" xml:"deleted"`

	// 帐号ID
	DomainId string `json:"domain_id" xml:"domain_id"`

	// 镜像是否被扫描过
	Scanned bool `json:"scanned" xml:"scanned"`

	// 0：manifest类型；1：manifest list类型
	TagType int64 `json:"tag_type" xml:"tag_type"`
}

func (o ShowReposTagResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReposTagResp struct{}"
	}

	return strings.Join([]string{"ShowReposTagResp", string(data)}, " ")
}
