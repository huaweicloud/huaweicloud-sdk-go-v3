package model

import (
	"encoding/json"

	"strings"
)

type CreateRequestBody struct {
	// BCS服务名

	Name string `json:"name"`
	// BCS服务版本类型，可选：基础版（1），专业版（4），企业版（2），铂金版（3）

	VersionType *int32 `json:"version_type,omitempty"`
	// Fabric版本，可选：\"1.4\"，\"2.0\"。目前HCS只支持1.4

	FabricVersion *string `json:"fabric_version,omitempty"`
	// 区块链类型，可选：联盟链（union），私有链（private）

	BlockchainType *string `json:"blockchain_type,omitempty"`
	// BCS服务的共识策略，可选：（etcdraft,1.4版本不支持raft共识算法）、快速拜占庭容错算法（SFLIC）、测试策略（solo）、Kafka共识（kafka）

	Consensus *string `json:"consensus,omitempty"`
	// BCS服务安全机制，可选：ECDSA（ECDSA），国密算法（sm2）

	SignAlgorithm *string `json:"sign_algorithm,omitempty"`
	// BCS服务所属企业项目ID

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// CCE集群存储卷类型，根据实际环境可选：云硬盘存储卷（evs），文件存储卷（nfs）, 极速文件存储卷（efs）

	VolumeType *string `json:"volume_type,omitempty"`
	// 云硬盘存储卷类型，volume_type选择evs时必填，可选：普通I/O（SATA），高I/O（SAS），超高I/O（SSD）

	EvsDiskType *string `json:"evs_disk_type,omitempty"`
	// [节点组织存储容量，基础版至少40GB，专业版和企业版至少100GB，铂金版至少500GB](tag:online)[节点组织存储容量GB，至少为100GB](tag:hcs)

	OrgDiskSize *int32 `json:"org_disk_size,omitempty"`
	// BCS服务数据库类型，包括文件数据库（goleveldb），NoSQL（couchdb）

	DatabaseType *string `json:"database_type,omitempty"`
	// BCS服务资源、区块链管理密码

	ResourcePassword string `json:"resource_password"`
	// 共识组织节点数，被邀请方创实例时可不填

	OrdererNodeNumber *int32 `json:"orderer_node_number,omitempty"`
	// 是否使用集群节点弹性IP

	UseEip *bool `json:"use_eip,omitempty"`
	// 弹性IP带宽

	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`
	// 集群类型，[可选：CCE集群（cce），边缘集群（ief）](tag:online)[目前线下混合云模式下只支持CCE集群](tag:hcs)

	ClusterType *string `json:"cluster_type,omitempty"`
	// 是否创建新集群

	CreateNewCluster bool `json:"create_new_cluster"`

	CceClusterInfo *CreateRequestBodyCceClusterInfo `json:"cce_cluster_info"`

	CceCreateInfo *CreateRequestBodyCceCreateInfo `json:"cce_create_info"`
	// IEF集群部署方式，随机部署（0），组织节点绑定（1）

	IefDeployMode *int32 `json:"ief_deploy_mode,omitempty"`
	// IEF集群节点列表

	IefNodesInfo *[]IefNode `json:"ief_nodes_info,omitempty"`
	// 节点组织列表

	PeerOrgs *[]OrgPeer `json:"peer_orgs,omitempty"`
	// 通道列表

	Channels *[]ChannelInfoV2 `json:"channels,omitempty"`

	CouchdbInfo *CreateRequestBodyCouchdbInfo `json:"couchdb_info,omitempty"`

	TurboInfo *CreateRequestBodyTurboInfo `json:"turbo_info,omitempty"`

	BlockInfo *CreateRequestBodyBlockInfo `json:"block_info,omitempty"`

	KafkaCreateInfo *CreateRequestBodyKafkaCreateInfo `json:"kafka_create_info,omitempty"`
	// 是否添加可信计算平台

	Tc3Need *bool `json:"tc3_need,omitempty"`
	// 是否添加restful API支持

	RestfulApiSupport *bool `json:"restful_api_support,omitempty"`
	// 是否是被邀请方创建实例

	IsInvitee *bool `json:"is_invitee,omitempty"`

	InvitorInfos *CreateRequestBodyInvitorInfos `json:"invitor_infos,omitempty"`
}

func (o CreateRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRequestBody", string(data)}, " ")
}
