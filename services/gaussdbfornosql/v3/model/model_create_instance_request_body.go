package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRequestBody **参数解释：** '实例信息。' **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type CreateInstanceRequestBody struct {

	// **参数解释：** 实例名称，允许和已有名称重复。 **约束限制：** 区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。 **取值范围：** 长度为4~64位，必须以字母开头（A~Z或a~z）。 **默认取值：** 不涉及。
	Name string `json:"name"`

	Datastore *DatastoreOption `json:"datastore"`

	// **参数解释：** 区域ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 非空。
	Region string `json:"region"`

	// **参数解释：** 可用区ID。 **约束限制：** 不涉及。 **取值范围：** 取值：请参见查询所有实例规格信息中返回的“az_status” ，支持创建3可用区实例，中间以逗号隔开。 **默认取值：** 不涉及。
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释：** 虚拟私有云ID。 **约束限制：**  获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询VPC列表。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释：** 子网的网络ID。 **约束限制：** 不涉及。 **取值范围：**  获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询子网列表。 **默认取值：** 不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释：** 安全组ID。 **约束限制：** 不涉及。 **取值范围：**  获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询安全组列表。 **默认取值：** 不涉及。
	SecurityGroupId string `json:"security_group_id"`

	// **参数解释：** 数据库密码。 **约束限制：** 不涉及。 **取值范围：** 长度为8~32位，必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。 **默认取值：** 不涉及。
	Password string `json:"password"`

	// **参数解释：** 实例类型。 **约束限制：** 不涉及。 **取值范围：**   -  GeminiDB Cassandra支持经典部署模式集群类型，取值为“Cluster”。   -  GeminiDB Cassandra支持云原生部署模式集群类型，取值“CloudNativeCluster”。   -  GeminiDB Mongo4.0版本支持副本集类型，取值为“ReplicaSet”。   -  GeminiDB Influx支持经典部署模式集群类型，取值为“Cluster”。   -  GeminiDB Influx支持云原生部署模式集群增强类型，取值为“CloudNativeCluster”。   -  GeminiDB Influx支持经典部署模式单节点类型，取值为“InfluxdbSingle”。   -  GeminiDB Influx支持经典部署模式集群增强类型，取值为“EnhancedCluster”。   -  GeminiDB Redis支持经典部署模式Proxy集群类型，取值为“Cluster”。   -  GeminiDB redis支持云原生部署模式集群类型，取值为“CloudNativeCluster”。   -  GeminiDB Redis支持经典部署模式Cluster集群类型，取值为“RedisCluster”   -  GeminiDB Redis支持经典部署模式主备类型，取值为“Replication”。   -  GeminiDB兼容DynamoDB支持经典部署模式集群类型，取值为“Cluster”。   -  GeminiDB兼容DynamoDB支持云原生部署模式集群类型，取值“CloudNativeCluster”。   -  GeminiDB HBase支持经典部署模式集群类型，取值为“Cluster”。   -  GeminiDB HBase支持云原生部署模式集群类型，取值“CloudNativeCluster”。 **默认取值：** 不涉及。
	Mode string `json:"mode"`

	// **参数解释：** 产品类型。 **约束限制：** 当创建GeminiDB Redis云原生部署模式集群类型必传此参数。 **取值范围：** - Capacity 容量型 **默认取值：** 不涉及。
	ProductType *string `json:"product_type,omitempty"`

	// **参数解释：** 实例规格详情。 **约束限制：** 获取方法请参见查询所有实例规格信息中响应“flavors”字段下参数的值。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Flavor []CreateInstanceFlavorOption `json:"flavor"`

	// **参数解释：** 磁盘加密时的密钥ID，创建实例时不传该参数，表示不进行磁盘加密。 **约束限制：** 备份恢复新实例时不支持指定该参数，如果原实例使用了磁盘加密，新实例复用原实例加密的密钥ID。 该参数仅GeminiDB influx 云原生部署模式集群增强版实例类型、GeminiDB Cassandra经典部署模式实例类型、GeminiDB兼容DynamoDB实例类型支持。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	// **参数解释：** '参数模板ID。' **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ConfigurationId *string `json:"configuration_id,omitempty"`

	BackupStrategy *BackupStrategyOption `json:"backup_strategy,omitempty"`

	// **参数解释：** '企业项目ID。' **约束限制：**   -对于未开通企业多项目服务的用户，不传该参数。   -对于已开通企业多项目服务的用户，不传该参数时，表示为default企业项目。获取方式请参见《企业管理API参考》的“查询企业项目列表”响应消息表“enterprise_project字段数据结构说明”的“id”。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** '专属资源ID。' **约束限制：** 只有开通专属资源池后才可以下发此参数。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	// **参数解释：** SSL开关选项。 **约束限制：** 不涉及。 **取值范围：** - 取“0”，表示DDS实例默认不启用SSL连接。 - 取“1”，表示DDS实例默认启用SSL连接。 - 不传该参数时，默认不启用SSL连接。 **默认取值：** 不涉及。
	SslOption *string `json:"ssl_option,omitempty"`

	ChargeInfo *ChargeInfoOption `json:"charge_info,omitempty"`

	RestoreInfo *RestoreInfo `json:"restore_info,omitempty"`

	// **参数解释：** 数据库访问端口号。 **约束限制：** 目前仅支持GeminiDB Redis实例支持自定义端口，取值范围为：1024~65535，禁用端口号为：2180、2887、3887、6377、6378、6380、8018、8079、8091、8479、8484、8999、12017、12333、50069。  不指定端口时，创建GeminiDB Redis实例的访问端口默认为6379。  如果该实例计划用于搭建双活容灾场景，请配置为8635端口。 **取值范围：** 取值范围为：1024~65535，禁用端口号为：2180、2887、3887、6377、6378、6380、8018、8079、8091、8479、8484、8999、12017、12333、50069。 **默认取值：** 6379。
	Port *string `json:"port,omitempty"`

	AvailabilityZoneDetail *AvailabilityZoneDetail `json:"availability_zone_detail,omitempty"`

	LbAccessControlSettings *LbAccessControlSettings `json:"lb_access_control_settings,omitempty"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}
