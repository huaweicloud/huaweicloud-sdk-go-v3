package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Cluster struct {

	// 集群ID。
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId"`

	// 集群名称。
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName"`

	// 集群部署的节点总数。
	TotalNodeNum *string `json:"totalNodeNum,omitempty" xml:"totalNodeNum"`

	// 集群状态，包括： - starting：启动中的集群。 - running：运行中的集群。 - terminated：已删除的集群。 - failed：失败的集群。 - abnormal：异常的集群。 - terminating：删除中的集群。 - frozen：已冻结的集群。 - scaling-out：扩容中的集群。 - scaling-in：缩容中的集群。
	ClusterState *string `json:"clusterState,omitempty" xml:"clusterState"`

	// 集群进度描述。  安装集群进度包括： - Verifying cluster parameters：校验集群参数中 - Applying for cluster resources：申请集群资源中 - Creating VM：创建虚拟机中 - Initializing VM：初始化虚拟机中 - Installing MRS Manager：安装MRS Manager中 - Deploying cluster：部署集群中 - Cluster installation failed：集群安装失败  扩容集群进度包括： - Preparing for cluster expansion：准备扩容中 - Creating VM：创建虚拟机中 - Initializing VM：初始化虚拟机中 - Adding node to the cluster：节点加入集群中 - Cluster expansion failed：集群扩容失败  缩容集群进度包括： - Preparing for cluster shrink：正在准备缩容 - Decommissioning instance：实例退服中 - Deleting VM：删除虚拟机中 - Deleting node from the cluster：从集群删除节点中 - Cluster shrink failed：集群缩容失败 集群安装、扩容、缩容失败，stageDesc会显示失败的原因。
	StageDesc *string `json:"stageDesc,omitempty" xml:"stageDesc"`

	// 集群创建时间，十位时间戳。
	CreateAt *string `json:"createAt,omitempty" xml:"createAt"`

	// 集群更新时间，十位时间戳。
	UpdateAt *string `json:"updateAt,omitempty" xml:"updateAt"`

	// 开始计费时间。
	ChargingStartTime *string `json:"chargingStartTime,omitempty" xml:"chargingStartTime"`

	// 集群计费模式。
	BillingType *string `json:"billingType,omitempty" xml:"billingType"`

	// 集群工作区域。
	DataCenter *string `json:"dataCenter,omitempty" xml:"dataCenter"`

	// VPC名称。
	Vpc *string `json:"vpc,omitempty" xml:"vpc"`

	// VPC ID。
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId"`

	// 集群购买时长。
	Duration *string `json:"duration,omitempty" xml:"duration"`

	// 创建集群所需费用，系统自动计算。
	Fee *string `json:"fee,omitempty" xml:"fee"`

	// Hadoop组件版本信息。
	HadoopVersion *string `json:"hadoopVersion,omitempty" xml:"hadoopVersion"`

	// 组件列表信息。
	ComponentList *[]ComponentAmb `json:"componentList,omitempty" xml:"componentList"`

	// 公网IP地址。
	ExternalIp *string `json:"externalIp,omitempty" xml:"externalIp"`

	// 公网备用IP地址。
	ExternalAlternateIp *string `json:"externalAlternateIp,omitempty" xml:"externalAlternateIp"`

	// 内网IP地址。
	InternalIp *string `json:"internalIp,omitempty" xml:"internalIp"`

	// 集群部署ID。
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId"`

	// 集群备注信息。
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 创建集群的订单号。
	OrderId *string `json:"orderId,omitempty" xml:"orderId"`

	// 可用区域ID。
	AzId *string `json:"azId,omitempty" xml:"azId"`

	// 可用区域名称。
	AzName *string `json:"azName,omitempty" xml:"azName"`

	// 可用区域英文名称
	AzCode *string `json:"azCode,omitempty" xml:"azCode"`

	// 可用区域
	AvailabilityZoneId *string `json:"availabilityZoneId,omitempty" xml:"availabilityZoneId"`

	// 实例ID。
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId"`

	// 远程登录弹性云服务器的URI地址。
	Vnc *string `json:"vnc,omitempty" xml:"vnc"`

	// 项目编号。
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId"`

	// 磁盘存储空间。
	VolumeSize *int32 `json:"volumeSize,omitempty" xml:"volumeSize"`

	// 磁盘类型。
	VolumeType *string `json:"volumeType,omitempty" xml:"volumeType"`

	// 子网ID。
	SubnetId *string `json:"subnetId,omitempty" xml:"subnetId"`

	// 子网名称。
	SubnetName *string `json:"subnetName,omitempty" xml:"subnetName"`

	// 安全组ID。
	SecurityGroupsId *string `json:"securityGroupsId,omitempty" xml:"securityGroupsId"`

	// 非Master节点的安全组id，当前一个MRS集群只会使用一个安全组，所以该字段已经废弃，从兼容性考虑，该字段会返回和securityGroupsId同样的值。
	SlaveSecurityGroupsId *string `json:"slaveSecurityGroupsId,omitempty" xml:"slaveSecurityGroupsId"`

	// 配置引导操作脚本信息。
	BootstrapScripts *[]BootstrapScriptResp `json:"bootstrapScripts,omitempty" xml:"bootstrapScripts"`

	// MRS集群运行模式。 - 0：普通集群 - 1：安全集群
	SafeMode *int32 `json:"safeMode,omitempty" xml:"safeMode"`

	// 集群版本。
	ClusterVersion *string `json:"clusterVersion,omitempty" xml:"clusterVersion"`

	// 密钥文件名称。
	NodePublicCertName *string `json:"nodePublicCertName,omitempty" xml:"nodePublicCertName"`

	// Master节点IP。
	MasterNodeIp *string `json:"masterNodeIp,omitempty" xml:"masterNodeIp"`

	// 首选私有IP。
	PrivateIpFirst *string `json:"privateIpFirst,omitempty" xml:"privateIpFirst"`

	// 错误信息。
	ErrorInfo *string `json:"errorInfo,omitempty" xml:"errorInfo"`

	// 标签信息
	Tags *string `json:"tags,omitempty" xml:"tags"`

	// 集群部署的Master节点数量。
	MasterNodeNum *string `json:"masterNodeNum,omitempty" xml:"masterNodeNum"`

	// 集群部署的Core节点数量。
	CoreNodeNum *string `json:"coreNodeNum,omitempty" xml:"coreNodeNum"`

	// Master节点的实例规格。
	MasterNodeSize *string `json:"masterNodeSize,omitempty" xml:"masterNodeSize"`

	// Core节点的实例规格。
	CoreNodeSize *string `json:"coreNodeSize,omitempty" xml:"coreNodeSize"`

	// Master节点产品ID。
	MasterNodeProductId *string `json:"masterNodeProductId,omitempty" xml:"masterNodeProductId"`

	// Master节点规格ID。
	MasterNodeSpecId *string `json:"masterNodeSpecId,omitempty" xml:"masterNodeSpecId"`

	// Core节点产品ID。
	CoreNodeProductId *string `json:"coreNodeProductId,omitempty" xml:"coreNodeProductId"`

	// Core节点规格ID。
	CoreNodeSpecId *string `json:"coreNodeSpecId,omitempty" xml:"coreNodeSpecId"`

	// Master节点数据磁盘存储类别，目前支持SATA、SAS和SSD。
	MasterDataVolumeType *string `json:"masterDataVolumeType,omitempty" xml:"masterDataVolumeType"`

	// Master节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。 取值范围：100GB～32000GB，传值只需填数字，不需要带单位GB。
	MasterDataVolumeSize *int32 `json:"masterDataVolumeSize,omitempty" xml:"masterDataVolumeSize"`

	// Master节点数据磁盘个数。 取值只能是1
	MasterDataVolumeCount *int32 `json:"masterDataVolumeCount,omitempty" xml:"masterDataVolumeCount"`

	// Core节点数据磁盘存储类别，目前支持SATA、SAS和SSD。
	CoreDataVolumeType *string `json:"coreDataVolumeType,omitempty" xml:"coreDataVolumeType"`

	// Core节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。 取值范围：100GB～32000GB，传值只需填数字，不需要带单位GB。
	CoreDataVolumeSize *int32 `json:"coreDataVolumeSize,omitempty" xml:"coreDataVolumeSize"`

	// Core节点数据磁盘个数。 取值范围：1～10
	CoreDataVolumeCount *int32 `json:"coreDataVolumeCount,omitempty" xml:"coreDataVolumeCount"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty" xml:"enterpriseProjectId"`

	// 表示集群创建过程中，MRS Manager是否安装完成。 - true：安装完成 - false：安装未完成
	IsMrsManagerFinish *bool `json:"isMrsManagerFinish,omitempty" xml:"isMrsManagerFinish"`

	// 集群类型。
	ClusterType *int32 `json:"clusterType,omitempty" xml:"clusterType"`

	// 集群安装失败时，是否搜集日志。 - 0：不收集 - 1：收集
	LogCollection *int32 `json:"logCollection,omitempty" xml:"logCollection"`

	// 区分包周期，集群是包年还是包月。 - 0：包月 - 1：包年
	PeriodType *int32 `json:"periodType,omitempty" xml:"periodType"`

	// 集群节点的变更状态（扩容/缩容/变更规格）。当该参数取值为空时，表示集群节点没有进行变更操作。 取值范围： - scaling-out：扩容中 - scaling-in：缩容中 - scaling-error：处于running状态，且上一次扩容/缩容/升级规格失败的集群 - scaling-up：Master节点规格升级中 - scaling_up_first：备Master节点规格升级中 - scaled_up_first：备Master节点规格升级成功 - scaled-up-success：Master节点规格升级成功
	Scale *string `json:"scale,omitempty" xml:"scale"`

	// Master节点、Core节点和Task节点列表信息。
	NodeGroups *[]NodeGroupV10 `json:"nodeGroups,omitempty" xml:"nodeGroups"`

	// Task节点列表信息。
	TaskNodeGroups *[]NodeGroupV10 `json:"taskNodeGroups,omitempty" xml:"taskNodeGroups"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}
