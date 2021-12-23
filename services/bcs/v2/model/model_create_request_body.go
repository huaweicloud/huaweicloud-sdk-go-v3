package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRequestBody struct {
	// BCS服务名

	Name string `json:"name"`
	// BCS服务版本类型，[专业版（4），企业版（2）](tag:g42)[可选：基础版（1），专业版（4），企业版（2），铂金版（3）](tag:online)。被邀请方创建时，和邀请方保持一致。

	VersionType int64 `json:"version_type"`
	// Fabric版本，当前邀请方以及私有链的创建仅可选：2.2 ;被邀请方创建时，和邀请方保持一致，1.4版本服务仅支持1.15及以下版本集群

	FabricVersion string `json:"fabric_version"`
	// 区块链类型，默认私有链，可选：联盟链（union），私有链（private）。被邀请方创建时，和邀请方保持一致。

	BlockchainType *string `json:"blockchain_type,omitempty"`
	// BCS服务的共识策略，Fabric1.4版本可选：测试策略（solo）、快速拜占庭容错算法（SFLIC）；Fabric2.2版本可选：raft共识算法（etcdraft）、快速拜占庭容错算法（SFLIC）。被邀请方创建时，和邀请方保持一致。

	Consensus *string `json:"consensus,omitempty"`
	// BCS服务安全机制，可选：ECDSA（ECDSA），国密算法（sm2）

	SignAlgorithm *string `json:"sign_algorithm,omitempty"`
	// BCS服务所属企业项目ID

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// CCE集群存储卷类型，根据实际环境可选：云硬盘存储卷（evs），文件存储卷（nfs）, 极速文件存储卷（efs）

	VolumeType *string `json:"volume_type,omitempty"`
	// 云硬盘存储卷类型，volume_type选择evs时必填，可选：普通I/O（SATA），高I/O（SAS），超高I/O（SSD）

	EvsDiskType *string `json:"evs_disk_type,omitempty"`
	// 节点组织存储容量[，基础版至少40GB，专业版和企业版至少100GB，铂金版至少500GB](tag:online)[，专业版和企业版至少100GB](tag:g42)[节点组织存储容量GB，至少为100GB](tag:hcs)

	OrgDiskSize *int64 `json:"org_disk_size,omitempty"`
	// BCS服务数据库类型，包括文件数据库（goleveldb），NoSQL（couchdb），选择couchdb需要填写couchdb_info字段中的信息

	DatabaseType *string `json:"database_type,omitempty"`
	// BCS服务资源、区块链管理密码

	ResourcePassword string `json:"resource_password"`
	// 共识组织节点数，被邀请方创实例时可不填。购买fabric2.2服务时必填。

	OrdererNodeNumber *int64 `json:"orderer_node_number,omitempty"`
	// 是否使用集群节点弹性IP

	UseEip *bool `json:"use_eip,omitempty"`
	// 弹性IP带宽

	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`
	// 集群类型，可选：CCE集群 [,边缘集群ief](tag:hasief)。

	ClusterType string `json:"cluster_type"`
	// 是否创建新集群，使用已有集群需要填写cce_cluster_info字段中的信息，创建新集群需要填写cce_create_info字段中的信息

	CreateNewCluster bool `json:"create_new_cluster"`

	CceClusterInfo *CceClusterInfo `json:"cce_cluster_info,omitempty"`

	CceCreateInfo *CceCreateInfo `json:"cce_create_info,omitempty"`
	// IEF集群部署方式，随机部署（0），组织节点绑定（1）。组织节点绑定模式时，peer_orgs 参数必填。组织名和IEF节点名必须一致。

	IefDeployMode *int64 `json:"ief_deploy_mode,omitempty"`
	// IEF集群节点列表，使用边缘集群模式部署时必填。

	IefNodesInfo *[]IefNode `json:"ief_nodes_info,omitempty"`
	// 节点组织列表。节点绑定模式中，组织名和IEF节点名必须一致。边缘集群模式时此字段必填。

	PeerOrgs *[]OrgPeer `json:"peer_orgs,omitempty"`
	// 通道列表

	Channels *[]ChannelInfoV2 `json:"channels,omitempty"`

	CouchdbInfo *Couchdb `json:"couchdb_info,omitempty"`

	TurboInfo *TurboInfo `json:"turbo_info,omitempty"`

	BlockInfo *CreateRequestBodyBlockInfo `json:"block_info,omitempty"`

	KafkaCreateInfo *KafkaCreateInfo `json:"kafka_create_info,omitempty"`
	// 是否添加可信计算平台

	Tc3Need *bool `json:"tc3_need,omitempty"`
	// 是否添加restful API支持

	RestfulApiSupport *bool `json:"restful_api_support,omitempty"`
	// 是否是创建被邀请方实例，创建被邀请方实例需要同时填写invitor_infos字段中的信息

	IsInvitee *bool `json:"is_invitee,omitempty"`

	InvitorInfos *InvitorInfos `json:"invitor_infos,omitempty"`
}

func (o CreateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRequestBody", string(data)}, " ")
}
