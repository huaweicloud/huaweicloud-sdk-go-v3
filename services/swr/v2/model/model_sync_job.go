package model

import (
	"encoding/json"

	"strings"
)

type SyncJob struct {
	// 创建时间，UTC日期格式，时间为UTC标准时间，用户需要根据本地时间计算偏移量；如东8区需要+8:00

	CreatedAt string `json:"createdAt"`
	// 租户ID

	DomainID string `json:"domainID"`
	// 租户名

	DomainName string `json:"domainName"`
	// ID

	Id int32 `json:"id"`
	// 组织名

	Namespace string `json:"namespace"`
	// 是否覆盖

	Override bool `json:"override"`
	// 目的组织

	RemoteNamespace string `json:"remoteNamespace"`
	// 目的region

	RemoteRegionId string `json:"remoteRegionId"`
	// 仓库名

	RepoName string `json:"repoName"`
	// 同步状态,waiting、running、success、failed、timeout、cancel、existed

	Status string `json:"status"`
	// 操作用户ID

	SyncOperatorId string `json:"syncOperatorId"`
	// 操作用户名

	SyncOperatorName string `json:"syncOperatorName"`
	// 镜像版本

	Tag string `json:"tag"`
	// updatedAt

	UpdatedAt string `json:"updatedAt"`
}

func (o SyncJob) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SyncJob struct{}"
	}

	return strings.Join([]string{"SyncJob", string(data)}, " ")
}
