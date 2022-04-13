package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicInfo struct {
	// BCS服务ID

	Id *string `json:"id,omitempty"`
	// BCS服务名

	Name *string `json:"name,omitempty"`
	// BCS服务版本信息

	Version *string `json:"version,omitempty"`
	// BCS服务的类型，分为：联盟链（union），私有链（private）

	ServiceType *string `json:"service_type,omitempty"`
	// BCS服务部署类型，一键部署（onestep），普通部署（normal）

	PurchaseType *string `json:"purchase_type,omitempty"`
	// BCS服务安全机制，分为ECDSA（ECDSA），国密算法（sm2）

	SignAlgorithm *string `json:"sign_algorithm,omitempty"`
	// BCS服务的共识策略，分为测试策略（solo），快速拜占庭容错算法（sflic）,Kafka(kafka)，raft共识算法（etcdraft）

	Consensus *string `json:"consensus,omitempty"`
	// BCS服务付费模式，分为按需（1）[包周期（0）](tag:onorder)

	ChargingMode *int64 `json:"charging_mode,omitempty"`
	// BCS服务版本类型

	VersionType *int64 `json:"version_type,omitempty"`
	// BCS服务数据库类型，包括文件数据库（goleveldb），NoSQL（couchdb）

	DatabaseType *string `json:"database_type,omitempty"`
	// BCS服务所在集群ID

	ClusterId *string `json:"cluster_id,omitempty"`
	// BCS服务所在集群名称

	ClusterName *string `json:"cluster_name,omitempty"`
	// BCS服务的集群类型，分为CCE集群（CCE），IEF集群（ief）

	ClusterType *string `json:"cluster_type,omitempty"`
	// BCS多可用区标示，分为：多可用区（yes），非多可用区（no）

	ClusterAz *string `json:"cluster_az,omitempty"`
	// BCS服务创建时间

	CreatedTime *string `json:"created_time,omitempty"`
	// BCS服务联盟链下生效，分为邀请方（create），被邀请方（invite）

	DeployType *string `json:"deploy_type,omitempty"`
	// BCS服务是否跨region

	IsCrossRegion *bool `json:"is_cross_region,omitempty"`
	// BCS服务升级失败，是否支持回滚

	IsSupportRollback *bool `json:"is_support_rollback,omitempty"`
	// BCS服务是否添加RESTful APIs支持，分为支持（true），不支持（false）

	IsSupportRestful *bool `json:"is_support_restful,omitempty"`
	// 区分BCS是否新服务，分为老服务（true），新服务（false）

	IsOldService *bool `json:"is_old_service,omitempty"`
	// BCS服务为老服务时，此字段为老服务版本号

	OldServiceVersion *string `json:"old_service_version,omitempty"`
	// BCS服务用户数据面agent地址端口列表

	AgentPortalAddrs *[]string `json:"agent_portal_addrs,omitempty"`
	// BCS服务状态，分为正常（Normal），异常（Abnormal），弹性IP异常（EipAbnormal），已冻结（Freeze），休眠中（Hibernation），未知（其余值）

	Status *string `json:"status,omitempty"`
	// BCS服务处理状态，分为创建中（IsCreating），升级中（IsUpgrading），扩缩容中（IsScaling），删除中（IsDeleting），添加中（IsAdding）

	ProcessStatus *string `json:"process_status,omitempty"`
	// BCS服务为包周期模式时，返回值为0（订单未成功）,1（订单异常）,2（订单正常）

	OrderStatus *int64 `json:"order_status,omitempty"`
}

func (o BasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicInfo struct{}"
	}

	return strings.Join([]string{"BasicInfo", string(data)}, " ")
}
