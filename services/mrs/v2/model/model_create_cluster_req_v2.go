package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type CreateClusterReqV2 struct {
	// 说明是否为专属云的资源，默认为false。

	IsDecProject *bool `json:"is_dec_project,omitempty"`
	// 集群版本。 MRS目前支持MRS 1.9.2、MRS 3.1.0版本。

	ClusterVersion string `json:"cluster_version"`
	// 集群名称，不允许相同。 只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。

	ClusterName string `json:"cluster_name"`
	// 集群类型，取值范围： - ANALYSIS：分析集群 - STREAMING：流式集群 - MIXED：混合集群 - CUSTOM：自定义集群，仅MRS 3.x版本支持。

	ClusterType string `json:"cluster_type"`

	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`
	// 集群所在区域信息，请参见[终端节点](https://support.huaweicloud.com/api-mrs/mrs_02_0003.html)。

	Region string `json:"region"`
	// 子网所在VPC名称。 通过VPC管理控制台获取名称： 1) 登录VPC管理控制台。 2) 单击“虚拟私有云”，从左侧列表选择虚拟私有云。  在“虚拟私有云”页面的列表中即可获取VPC名称。

	VpcName string `json:"vpc_name"`
	// 子网ID。通过VPC管理控制台获取子网ID： 1) 登录VPC管理控制台。 2) 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 3) 单击对应虚拟私有云所在行的“子网个数”查看子网。 4) 单击对应子网名称，获取“网络ID”。  “subnet_id”和“subnet_name”必须至少填写一个，当这两个参数同时配置但是不匹配同一个子网时，集群会创建失败，请仔细填写参数。推荐使用“subnet_id”。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 子网名称。 通过VPC管理控制台获取子网名称： 1) 登录管理控制台。 2) 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 3) 单击对应虚拟私有云所在行的“子网个数”查看子网，获取子网名称。  “subnet_id”和“subnet_name”必须至少填写一个，当这两个参数同时配置但是不匹配同一个子网时，集群会创建失败，请仔细填写参数。当仅填写“subnet_name”一个参数且VPC下存在同名子网时，创建集群时以VPC平台第一个名称的子网为准。推荐使用“subnet_id”。

	SubnetName string `json:"subnet_name"`
	// 组件名称列表，用逗号分隔。支持的组件请参见[获取MRS集群信息](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)页面的“MRS服务支持的组件”内容。

	Components string `json:"components"`
	// 可用分区名称。 可用分区信息请参见[终端节点](https://support.huaweicloud.com/api-mrs/mrs_02_0003.html)。

	AvailabilityZone string `json:"availability_zone"`
	// 集群安全组的ID。 - 当该ID为空时MRS后台会自动创建安全组，自动创建的安全组名称以mrs_{cluster_name}开头。 - 当该ID不为空时，表示使用固定安全组来创建集群，传入的ID必须是当前租户中包含的安全组ID，且该安全组中需要包含一条支持全部协议、全部端口、源地址为指定的管理面节点IP的入方向规则。

	SecurityGroupsId *string `json:"security_groups_id,omitempty"`
	// 是否要创建MRS集群默认安全组，默认为false。 当指定该参数为true，则无论“security_groups_id”参数是否指定，都会为集群创建默认安全组。

	AutoCreateDefaultSecurityGroup *bool `json:"auto_create_default_security_group,omitempty"`
	// MRS集群运行模式。 - SIMPLE：普通集群，表示Kerberos认证关闭，用户可使用集群提供的所有功能。 - KERBEROS：安全集群，表示Kerberos认证开启，普通用户无权限使用MRS集群的“文件管理”和“作业管理”功能，并且无法查看Hadoop、Spark的作业记录以及集群资源使用情况。如果需要使用集群更多功能，需要找Manager的管理员分配权限。

	SafeMode string `json:"safe_mode"`
	// 配置Manager管理员用户的密码。 - 密码长度应在8～32个字符之间 - 至少包含三种字符组合，如大写字母，小写字母，数字，特殊字符（`~!@#$%^&*()-_=+\\|[{}];:'\",<.>/?）和空格。 - 不能与用户名或者倒序用户名相同

	ManagerAdminPassword string `json:"manager_admin_password"`
	// 节点登录方式。 - PASSWORD：密码登录，选择此项时，node_root_password不能为空。 - KEYPAIR：密钥对登录，选择此项时，node_keypair_name不能为空。

	LoginMode string `json:"login_mode"`
	// 配置访问集群节点的root密码。 密码设置约束如下： - 字符串类型，可输入的字符串长度为8-26。 - 至少包含三种字符组合，如大写字母，小写字母，数字，特殊字符（!@$%^-_=+[{}]:,./?），但不能包含空格。 - 不能与用户名或者倒序用户名相同。

	NodeRootPassword *string `json:"node_root_password,omitempty"`
	// 密钥对名称。用户可以使用密钥对方式登录集群节点。

	NodeKeypairName *string `json:"node_keypair_name,omitempty"`
	// 企业项目ID。  创建集群时，给集群绑定企业项目ID。  默认设置为0，表示为default企业项目。  获取方式请参见《企业管理API参考》的“查询企业项目列表”响应消息表“enterprise_project字段数据结构说明”的“id”。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 与MRS集群绑定的弹性公网IP，可实现使用弹性公网IP访问Manager的目的。该弹性公网IP必须已经创建且与集群在同一区域。

	EipAddress *string `json:"eip_address,omitempty"`
	// 当“eip_address”配置时，该参数必须配置，用于表示绑定的弹性公网IP的ID。可通过在VPC服务的“网络控制台 > 弹性公网IP和带宽 > 弹性公网IP”页面单击待绑定的弹性公网IP，在基本信息中获取“ID”。

	EipId *string `json:"eip_id,omitempty"`
	// 集群节点默认绑定的委托名称，固定为MRS_ECS_DEFAULT_AGENCY。 通过绑定委托，您可以将部分资源共享给ECS或BMS云服务来管理，例如通过配置ECS委托可自动获取AK/SK访问OBS。 MRS_ECS_DEFAULT_AGENCY委托拥有对象存储服务的OBS OperateAccess权限和在集群所在区域拥有CES FullAccess（对开启细粒度策略的用户）、CES Administrator和KMS Administrator权限。

	MrsEcsDefaultAgency *string `json:"mrs_ecs_default_agency,omitempty"`
	// 当集群类型为CUSTOM时，用于指定节点部署所使用的模板。 - mgmt_control_combined_v2：管控合设模板，管理角色和控制角色共同部署在Master节点中，数据实例合设在同一节点组。该部署方式适用于100个以下的节点，可以减少成本。 - mgmt_control_separated_v2：管控分设模板，管理角色和控制角色分别部署在不同的Master节点中，数据实例合设在同一节点组。该部署方式适用于100-500个节点，在高并发负载情况下表现更好。 - mgmt_control_data_separated_v2：数据分设模板，管理角色和控制角色分别部署在不同的Master节点中，数据实例分设在不同节点组。该部署方式适用于500个以上的节点，可以将各组件进一步分开部署，适用于更大的集群规模。

	TemplateId *string `json:"template_id,omitempty"`
	// 集群的标签信息。 同一个集群最多能使用10个tag，tag的名称（key）不能重复。

	Tags *[]Tag `json:"tags,omitempty"`
	// 集群创建失败时，是否收集失败日志。 默认设置为1，此时将创建OBS桶仅用于MRS集群创建失败时的日志收集。  枚举值： - 0：不收集 - 1：收集

	LogCollection *CreateClusterReqV2LogCollection `json:"log_collection,omitempty"`
	// 组成集群的节点组信息。

	NodeGroups []NodeGroupV2 `json:"node_groups"`
	// 配置引导操作脚本信息。

	BootstrapScripts *[]BootstrapScript `json:"bootstrap_scripts,omitempty"`
	// 创建集群时可同时提交作业，当前版本暂时只支持新增一个作业。

	AddJobs *[]AddJobs `json:"add_jobs,omitempty"`
}

func (o CreateClusterReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReqV2 struct{}"
	}

	return strings.Join([]string{"CreateClusterReqV2", string(data)}, " ")
}

type CreateClusterReqV2LogCollection struct {
	value int32
}

type CreateClusterReqV2LogCollectionEnum struct {
	E_0 CreateClusterReqV2LogCollection
	E_1 CreateClusterReqV2LogCollection
}

func GetCreateClusterReqV2LogCollectionEnum() CreateClusterReqV2LogCollectionEnum {
	return CreateClusterReqV2LogCollectionEnum{
		E_0: CreateClusterReqV2LogCollection{
			value: 0,
		}, E_1: CreateClusterReqV2LogCollection{
			value: 1,
		},
	}
}

func (c CreateClusterReqV2LogCollection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV2LogCollection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
