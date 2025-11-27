package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterSpec struct {

	// 集群和karmada控制面之间的同步模式
	SyncMode *string `json:"syncMode,omitempty"`

	// 容器舰队id
	ClusterGroupID *string `json:"clusterGroupID,omitempty"`

	// 集群类型，取值如下： - grouped：在舰队中纳管的集群 - discrete：未加入舰队的集群
	ManageType *string `json:"manageType,omitempty"`

	// 集群包含的RuleNamespace列表
	RuleNamespaces *[]RuleNamespace `json:"ruleNamespaces,omitempty"`

	// apiserver地址
	ApiEndpoint *string `json:"apiEndpoint,omitempty"`

	SecretRef *LocalSecretReference `json:"secretRef,omitempty"`

	// 是否跳过https校验
	InsecureSkipTLSVerification *bool `json:"insecureSkipTLSVerification,omitempty"`

	// 代理链接
	ProxyURL *string `json:"proxyURL,omitempty"`

	// 提供商
	Provider *string `json:"provider,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 类别
	Category *string `json:"category,omitempty"`

	// CCE Turbo集群是否支持管理边缘基础设施。
	EnableDistMgt *bool `json:"enableDistMgt,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 国家
	Country *string `json:"country,omitempty"`

	// 城市
	City *string `json:"city,omitempty"`

	// 项目id
	ProjectID *string `json:"projectID,omitempty"`

	// 项目名
	ProjectName *string `json:"projectName,omitempty"`

	// 地区
	Zone *string `json:"zone,omitempty"`

	// 污点
	Taints *[]Taint `json:"taints,omitempty"`

	// 是否已经下载过证书
	IsDownloadedCert *bool `json:"IsDownloadedCert,omitempty"`

	// 策略管理id
	PolicyId *string `json:"policyId,omitempty"`

	// 集群所属domainID
	OperatorNamespace *string `json:"operatorNamespace,omitempty"`

	// 连接代理服务器的终端节点服务ID列表，仅非华为云集群返回该字段
	ConnectProxyEndpoints *[]ConnectEndpoint `json:"connectProxyEndpoints,omitempty"`
}

func (o ClusterSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSpec struct{}"
	}

	return strings.Join([]string{"ClusterSpec", string(data)}, " ")
}
