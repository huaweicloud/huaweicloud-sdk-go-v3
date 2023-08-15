package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpengaussRestoreInstanceRequest 实例信息。
type OpengaussRestoreInstanceRequest struct {

	// 实例名称。 用于表示实例的名称，同一租户下，同类型的实例名可重名。 取值范围：4~64个字符之间，必须以字母开头，区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符。
	Name string `json:"name"`

	// 可用区ID。  GaussDB取值范围：非空，可选部署在同一可用区或三个不同可用区，可用区之间用逗号隔开。详见示例。  - 部署在同一可用区：需要输入三个相同的可用区。例如：部署在“cn-north-4a”可用区，则需要在此处输入\"cn-north-4a,cn-north-4a,cn-north-4a\"。 - 部署在三个不同可用区：需要分别输入三个不同的可用区。 取值范围：非空，请参见[地区和终端节点](https://developer.huaweicloud.com/endpoint)。
	AvailabilityZone string `json:"availability_zone"`

	// 规格码，取值范围：非空。参考[表1](https://support.huaweicloud.com/api-opengauss/opengauss_api_0037.html#opengauss_api_0037__ted9b9d433c8a4c52884e199e17f94479)中GaussDB 的“规格编码”列内容获取。
	FlavorRef string `json:"flavor_ref"`

	Volume *OpenGaussVolume `json:"volume"`

	// 用于磁盘加密的密钥ID，默认为空。
	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	// 虚拟私有云ID，获取方法如下：  - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询VPC列表](https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html)。
	VpcId string `json:"vpc_id"`

	// 子网的网络ID信息，获取方法如下：  - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。
	SubnetId string `json:"subnet_id"`

	// 指定实例所属的安全组。如果不需要指定安全组，请联系客服申请白名单。  - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询安全组列表](https://support.huaweicloud.com/api-vpc/vpc_sg01_0003.html)。
	SecurityGroupId string `json:"security_group_id"`

	// 数据库密码。  取值范围：  非空，由大小写字母、数字和特殊符号~!@#%^*-_=+?组成，长度8~32个字符。  建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。
	Password string `json:"password"`

	ChargeInfo *OpenGaussChargeInfo `json:"charge_info,omitempty"`

	RestorePoint *RestorePoint `json:"restore_point"`

	BackupStrategy *OpenGaussBackupStrategy `json:"backup_strategy,omitempty"`

	// 是否支持备份并行恢复。当不传该参数时，企业版默认为不支持，主备版默认支持。
	EnableParallelRestore *bool `json:"enable_parallel_restore,omitempty"`

	// 参数组ID，当不传该参数时，使用系统默认的参数模板。
	ConfigurationId *string `json:"configuration_id,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 数据库对外开放的端口，不填默认为8000，可选范围为：1024-39998。限制端口： 2378,2379,2380,4999,5000,5999,6000,6001,8097,8098,12016,12017,20049,20050,21731,21732,32122,32123,32124。  - GaussDB数据库端口当前只支持设置为8000，当不传该参数时，默认端口为8000。
	Port *string `json:"port,omitempty"`

	// 时区。  - 不选择时，国内站默认为UTC+08:00，国际站默认为UTC时间。 - 选择填写时，取值范围为UTC-12:00~UTC+12:00，且只支持整段时间，如UTC+08:00，不支持UTC+08:30。
	TimeZone *string `json:"time_zone,omitempty"`

	// enable_force_switch表示是否开启备机强升主功能，仅支持取值true，false。 enable_force_switch=true表示开启备机强升主功能，enable_force_switch=false表示关闭，默认关闭。仅支持1.2.2及以上版本。  说明：  备机强升主功能适用场景：在主机发生故障后，为了保障集群的可用性，强制拉起备机作为新主机对外提供服务的场景。 本功能在集群故障状态下，以丢失部分数据为代价换取集群尽可能快的恢复服务。本功能是集群状态为不可用时的一个逃生方法，如果操作者不清楚备机强升后丢失数据对业务的影响，请勿使用本功能。 备机强升主相关介绍请参考《故障处理》备机强升主章节。
	EnableForceSwitch *bool `json:"enable_force_switch,omitempty"`
}

func (o OpengaussRestoreInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpengaussRestoreInstanceRequest struct{}"
	}

	return strings.Join([]string{"OpengaussRestoreInstanceRequest", string(data)}, " ")
}
