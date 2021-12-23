package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例信息。
type CreateInstanceRequestBody struct {
	// 实例名称，允许和已有名称重复。 取值范围：长度为4~64位，必须以字母开头（A~Z或a~z），区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。

	Name string `json:"name"`

	Datastore *DatastoreOption `json:"datastore"`
	// - 区域ID - 取值：非空。

	Region string `json:"region"`
	// 可用区ID。 取值：请参见查询所有实例规格信息中返回的“az_status” ，支持创建3可用区实例，中间以逗号隔开。

	AvailabilityZone string `json:"availability_zone"`
	// 虚拟私有云ID，获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询VPC列表。

	VpcId string `json:"vpc_id"`
	// 子网的网络ID，获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询子网列表。

	SubnetId string `json:"subnet_id"`
	// 安全组ID，获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询安全组列表。

	SecurityGroupId string `json:"security_group_id"`
	// 数据库密码。 取值范围：长度为8~32位，必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。

	Password string `json:"password"`
	// 实例类型。   - GaussDB(for Cassandra)支持集群类型，取值为“Cluster”。   - GaussDB(for Mongo)3.4版本支持集群类型，取值为“Sharding”   - GaussDB(for Mongo)4.0版本支持副本集类型，取值为“ReplicaSet”。   - GaussDB(for Influx)支持集群类型，取值为“Cluster”。

	Mode string `json:"mode"`
	// 实例规格详情。获取方法请参见查询所有实例规格信息中响应“flavors”字段下参数的值。

	Flavor []CreateInstanceFlavorOption `json:"flavor"`
	// 参数模板ID。

	ConfigurationId *string `json:"configuration_id,omitempty"`

	BackupStrategy *BackupStrategyOption `json:"backup_strategy,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 专属资源ID，只有开通专属资源池后才可以下发此参数。

	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`
	// SSL开关选项。 取值： - 取“0”，表示DDS实例默认不启用SSL连接。 - 取“1”，表示DDS实例默认启用SSL连接。 - 不传该参数时，默认不启用SSL连接。

	SslOption *string `json:"ssl_option,omitempty"`

	ChargeInfo *ChargeInfoOption `json:"charge_info,omitempty"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}
