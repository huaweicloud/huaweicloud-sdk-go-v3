package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OnPremisesVersionInfo struct {

	// Kubernetes集群版本
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	// 发布版本
	ReleaseVersion *string `json:"releaseVersion,omitempty"`

	// ucs-ctl工具下载链接
	UcsctlDownloadURL *string `json:"ucsctlDownloadURL,omitempty"`

	// 兼容CPU架构
	Arch *string `json:"arch,omitempty"`

	// 对象存储服务访问端点
	ObsEndpoint *string `json:"obsEndpoint,omitempty"`

	// 软件包存储路径
	PackagePath *string `json:"packagePath,omitempty"`

	// 镜像包存储路径
	ImagesPackagePath *string `json:"imagesPackagePath,omitempty"`

	// 目标版本
	TargetVersion *string `json:"targetVersion,omitempty"`
}

func (o OnPremisesVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OnPremisesVersionInfo struct{}"
	}

	return strings.Join([]string{"OnPremisesVersionInfo", string(data)}, " ")
}
