package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCloudClusterCreateRequestBody 多云集群的创建
type MultiCloudClusterCreateRequestBody struct {

	// 集群名称
	Name *string `json:"name,omitempty"`

	// **参数解释** 集群服务商 **约束限制**   - ali：阿里。   - tencent：腾讯。   - azure：微软。   - aws：亚马逊。   - self_built_hw：华为自建。   - self_built_idc：IDC自建。  **取值范围** 字符长度范围0-64 **默认取值** 不涉及
	Provider *string `json:"provider,omitempty"`

	// 集群apiserver地址
	Server *string `json:"server,omitempty"`

	// 镜像仓地址
	ImageRepo *string `json:"image_repo,omitempty"`

	// 镜像仓用户名
	ImageRepoUsername *string `json:"image_repo_username,omitempty"`

	// 镜像仓密码
	ImageRepoPassword *string `json:"image_repo_password,omitempty"`

	// 组织
	Organization *string `json:"organization,omitempty"`

	// **参数解释** 镜像仓类型 **约束限制** - harbor：Harbor镜像仓。 - quay：Quay镜像仓。 - jfrog：Jfrog镜像仓。 - other：其他镜像仓。  **取值范围** 字符长度范围0-64 **默认取值** 不涉及
	ImageRepoType *string `json:"image_repo_type,omitempty"`

	// 当前有效期结束时间
	CurrentExpirationDate *int64 `json:"current_expiration_date,omitempty"`

	// 证书有效期结束时间
	CertificateExpirationDate *int64 `json:"certificate_expiration_date,omitempty"`

	// kubeconfig文件
	KubeConfig string `json:"kube_config"`

	// 镜像架构类型： - x86 - arm
	ImageArch *string `json:"image_arch,omitempty"`

	// Anp代理地址
	AnpProxy *string `json:"anp_proxy,omitempty"`

	// Hostguard代理地址
	HostguardProxy *string `json:"hostguard_proxy,omitempty"`

	// 容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// DNS策略： - default：继承集群节点的域名解析配置 - clusterfirst：查询集群内部的CoreDNS服务 - clusterfirstwithhostnet：强制在hostNetwork网络模式下使用ClusterFirst策略 - none：忽略集群的DNS策略，使用自定义DNS配置
	DnsPolicy *string `json:"dns_policy,omitempty"`

	// 自定义DNS配置，一个或多个IP地址，以分号分隔
	DnsConfig *string `json:"dns_config,omitempty"`
}

func (o MultiCloudClusterCreateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCloudClusterCreateRequestBody struct{}"
	}

	return strings.Join([]string{"MultiCloudClusterCreateRequestBody", string(data)}, " ")
}
