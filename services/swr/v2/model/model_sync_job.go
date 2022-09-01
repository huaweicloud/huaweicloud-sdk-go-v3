package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncJob struct {

	// 创建时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00
	CreatedAt string `json:"createdAt" xml:"createdAt"`

	// 租户ID
	DomainID string `json:"domainID" xml:"domainID"`

	// 租户名
	DomainName string `json:"domainName" xml:"domainName"`

	// ID
	Id int32 `json:"id" xml:"id"`

	// 组织名
	Namespace string `json:"namespace" xml:"namespace"`

	// 是否覆盖
	Override bool `json:"override" xml:"override"`

	// 目的组织
	RemoteNamespace string `json:"remoteNamespace" xml:"remoteNamespace"`

	// 目的region
	RemoteRegionId string `json:"remoteRegionId" xml:"remoteRegionId"`

	// 仓库名
	RepoName string `json:"repoName" xml:"repoName"`

	// 同步状态,waiting、running、success、failed、timeout、cancel、existed
	Status string `json:"status" xml:"status"`

	// 操作用户ID
	SyncOperatorId string `json:"syncOperatorId" xml:"syncOperatorId"`

	// 操作用户名
	SyncOperatorName string `json:"syncOperatorName" xml:"syncOperatorName"`

	// 镜像版本
	Tag string `json:"tag" xml:"tag"`

	// updatedAt
	UpdatedAt string `json:"updatedAt" xml:"updatedAt"`
}

func (o SyncJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncJob struct{}"
	}

	return strings.Join([]string{"SyncJob", string(data)}, " ")
}
