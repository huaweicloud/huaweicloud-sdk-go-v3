package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateClusterReqV11 struct {

	// 集群版本。 例如：MRS 3.1.0。
	ClusterVersion string `json:"cluster_version"`

	// 集群名称，不允许相同。 只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。
	ClusterName string `json:"cluster_name"`

	// Master节点数量。启用集群高可用功能时配置为2，不启用集群高可用功能时配置为1。MRS 3.x版本暂时不支持该参数配置为1。
	MasterNodeNum *int32 `json:"master_node_num,omitempty"`

	// Core节点数量。  取值范围：[1～500]  Core节点默认的最大值为500，如果用户需要的Core节点数大于500，请申请扩大配额。
	CoreNodeNum *int32 `json:"core_node_num,omitempty"`

	// 集群的计费模式。  12：表示按需计费。接口调用仅支持创建按需计费集群。
	BillingType CreateClusterReqV11BillingType `json:"billing_type"`

	// 集群区域信息，请参见[终端节点及区域](https://support.huaweicloud.com/api-mrs/mrs_02_0003.html)。
	DataCenter string `json:"data_center"`

	// 子网所在VPC名称。 通过VPC管理控制台获取名称： 1) 登录管理控制台。 2) 单击“虚拟私有云”，从左侧列表选择虚拟私有云。  在“虚拟私有云”页面的列表中即可获取VPC名称。
	Vpc string `json:"vpc"`

	// Master节点的实例规格，例如：c3.4xlarge.2.linux.bigdata。MRS当前支持主机规格的配型由CPU+内存+Disk共同决定。实例规格详细说明请参见[MRS所使用的弹性云服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9006.html)和[MRS所使用的裸金属服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9001.html)。 该参数建议从MRS控制台的集群创建页面获取对应区域对应版本所支持的规格。
	MasterNodeSize *string `json:"master_node_size,omitempty"`

	// Core节点的实例规格，例如：c3.4xlarge.2.linux.bigdata。实例规格详细说明请参见[MRS所使用的弹性云服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9006.html)和[MRS所使用的裸金属服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9001.html)。 该参数建议从MRS控制台的集群创建页面获取对应区域对应版本所支持的规格。
	CoreNodeSize *string `json:"core_node_size,omitempty"`

	// 服务组件安装列表信息。
	ComponentList []ComponentAmbV11 `json:"component_list"`

	// 可用分区ID。  - 华北-北京一可用区1（cn-north-1a）：ae04cf9d61544df3806a3feeb401b204 - 华北-北京一可用区2（cn-north-1b）：d573142f24894ef3bd3664de068b44b0 - 华东-上海二可用区1（cn-east-2a）：72d50cedc49846b9b42c21495f38d81c - 华东-上海二可用区2（cn-east-2b）：38b0f7a602344246bcb0da47b5d548e7 - 华东-上海二可用区3（cn-east-2c）：5547fd6bf8f84bb5a7f9db062ad3d015 - 华南-广州可用区1（cn-south-1a）：34f5ff4865cf4ed6b270f15382ebdec5 - 华南-广州可用区2（cn-south-2b）：043c7e39ecb347a08dc8fcb6c35a274e - 华南-广州可用区3（cn-south-1c）：af1687643e8c4ec1b34b688e4e3b8901 - 华北-北京四可用区1（cn-north-4a）：effdcbc7d4d64a02aa1fa26b42f56533 - 华北-北京四可用区2（cn-north-4b）：a0865121f83b41cbafce65930a22a6e8 - 华北-北京四可用区3（cn-north-4c）：2dcb154ac2724a6d92e9bcc859657c1e
	AvailableZoneId string `json:"available_zone_id"`

	// 子网所在VPC ID。 通过VPC管理控制台获取ID： 1) 登录管理控制台。 2) 单击“虚拟私有云”，从左侧列表选择虚拟私有云。   在“虚拟私有云”页面的列表中即可获取VPC ID。
	VpcId string `json:"vpc_id"`

	// 子网ID。通过VPC管理控制台获取子网ID： 1) 登录管理控制台。 2) 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 3) 单击对应虚拟私有云所在行的“子网个数”查看子网。 4) 单击对应子网名称，获取“网络ID”。  “subnet_id”和“subnet_name”必须至少填写一个，当这两个参数同时配置但是不匹配同一个子网时，集群会创建失败，请仔细填写参数。推荐使用“subnet_id”。
	SubnetId string `json:"subnet_id"`

	// 子网名称。 通过VPC管理控制台获取子网名称： 1) 登录管理控制台。 2) 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 3) 单击对应虚拟私有云所在行的“子网个数”查看子网，获取子网名称。  “subnet_id”和“subnet_name”必须至少填写一个，当这两个参数同时配置但是不匹配同一个子网时，集群会创建失败，请仔细填写参数。当仅填写“subnet_name”一个参数且VPC下存在同名子网时，创建集群时以VPC平台第一个名称的子网为准。推荐使用“subnet_id”。
	SubnetName string `json:"subnet_name"`

	// 集群安全组的ID。 - 当该ID为空时MRS后台会自己创建安全组，自动创建的安全组名称以mrs_{cluster_name}开头。 - 当该ID不为空时，表示使用固定安全组来创建集群，传入的ID必须是当前租户中包含的安全组ID，且该安全组中包含一条全部协议，全部端口，源地址为指定的管理面节点IP的入方向规则。
	SecurityGroupsId *string `json:"security_groups_id,omitempty"`

	// 创建集群时可同时提交作业，当前版本暂时只支持新增一个作业。
	AddJobs *[]AddJobsReqV11 `json:"add_jobs,omitempty"`

	// Master和Core节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。可以根据如下应用场景合理选择磁盘存储空间大小： - 数据存储和计算分离，数据存储在OBS系统中，集群费用相对较低，计算性能不高，并且集群随时可以删除，建议数据计算不频繁场景下使用。 - 数据存储和计算不分离，数据存储在HDFS中，集群费用相对较高，计算性能高，集群需要长期存在，建议数据计算频繁场景下使用。  取值范围：100GB～32000GB，传值只需填数字，不需要带单位GB。 不建议使用该参数，详情请参考volume_type参数的说明。
	VolumeSize *int32 `json:"volume_size,omitempty"`

	// Master和Core节点的磁盘存储类别，目前支持SATA、SAS、SSD和GPSSD。磁盘参数可以使用volume_type和volume_size表示，也可以使用多磁盘相关的参数表示。volume_type和volume_size这两个参数如果与多磁盘参数同时出现，系统优先读取volume_type和volume_size参数。建议使用多磁盘参数。 - SATA：普通IO - SAS：高IO - SSD：超高IO - GPSSD：通用型SSD
	VolumeType *CreateClusterReqV11VolumeType `json:"volume_type,omitempty"`

	// 该参数为多磁盘参数，表示Master节点数据磁盘存储类别，目前支持SATA、SAS、SSD和GPSSD。
	MasterDataVolumeType *CreateClusterReqV11MasterDataVolumeType `json:"master_data_volume_type,omitempty"`

	// 该参数为多磁盘参数，表示Master节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。  取值范围：100GB～32000GB，传值只需填数字，不需要带单位GB。
	MasterDataVolumeSize *int32 `json:"master_data_volume_size,omitempty"`

	// 该参数为多磁盘参数，表示Master节点数据磁盘个数。取值只能是1。
	MasterDataVolumeCount *CreateClusterReqV11MasterDataVolumeCount `json:"master_data_volume_count,omitempty"`

	// 该参数为多磁盘参数，表示Core节点数据磁盘存储类别，目前支持SATA、SAS、SSD和GPSSD。
	CoreDataVolumeType *CreateClusterReqV11CoreDataVolumeType `json:"core_data_volume_type,omitempty"`

	// 该参数为多磁盘参数，表示Core节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。  取值范围：100GB～32000GB，传值只需填数字，不需要带单位GB。
	CoreDataVolumeSize *int32 `json:"core_data_volume_size,omitempty"`

	// 该参数为多磁盘参数，表示Core节点数据磁盘个数。 取值范围：1～20
	CoreDataVolumeCount *int32 `json:"core_data_volume_count,omitempty"`

	// Task节点列表信息。
	TaskNodeGroups *[]TaskNodeGroup `json:"task_node_groups,omitempty"`

	// 配置引导操作脚本信息。
	BootstrapScripts *[]BootstrapScript `json:"bootstrap_scripts,omitempty"`

	// 密钥对名称。用户可以使用密钥对方式登录集群节点。当“login_mode”配置为“1”时，请求消息体中包含node_public_cert_name字段。
	NodePublicCertName *string `json:"node_public_cert_name,omitempty"`

	// 配置MRS Manager管理员用户的密码。 - 密码长度应在8～26个字符之间 - 不能与用户名或者倒序用户名相同 - 必须包含如下4种字符的组合     - 至少一个小写字母     - 至少一个大写字母     - 至少一个数字     - 至少一个特殊字符：!@$%^-_=+[{}]:,./?
	ClusterAdminSecret *string `json:"cluster_admin_secret,omitempty"`

	// 配置访问集群节点的root密码。当“login_mode”配置为“0”时，请求消息体中包含cluster_master_secret字段。  密码设置约束如下： - 字符串类型，可输入的字符串长度为8-26。 - 至少包含4种字符组合，如大写字母，小写字母，数字，特殊字符（!@$%^-_=+[{}]:,./?），但不能包含空格。 - 不能与用户名或者倒序用户名相同。
	ClusterMasterSecret *string `json:"cluster_master_secret,omitempty"`

	// MRS集群运行模式。 - 0：普通集群，表示Kerberos认证关闭，用户可使用集群提供的所有功能。 - 1：安全集群，表示Kerberos认证开启，普通用户无权限使用MRS集群的“文件管理”和“作业管理”功能，并且无法查看Hadoop、Spark的作业记录以及集群资源使用情况。如果需要使用集群更多功能，需要找MRS Manager的管理员分配权限。
	SafeMode CreateClusterReqV11SafeMode `json:"safe_mode"`

	// 集群类型。  默认值为0：分析集群。  说明：暂不支持通过接口方式创建混合集群。  枚举值： - 0：分析集群 - 1：流式集群
	ClusterType *CreateClusterReqV11ClusterType `json:"cluster_type,omitempty"`

	// 集群创建失败时，是否收集失败日志。  默认设置为1，将创建OBS桶仅用于MRS集群创建失败时的日志收集。  枚举值： - 0：不收集 - 1：收集
	LogCollection *CreateClusterReqV11LogCollection `json:"log_collection,omitempty"`

	// 企业项目ID。  创建集群时，给集群绑定企业项目ID。  默认设置为0，表示为default企业项目。  获取方式请参见《企业管理API参考》的“查询企业项目列表”响应消息表“enterprise_project字段数据结构说明”的“id”。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群的标签信息。  同一个集群最多能使用10个tag，tag的名称（key）不能重复 标签的键/值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。
	Tags *[]Tag `json:"tags,omitempty"`

	// 集群登录方式。默认设置为1。  - 当“login_mode”配置为“0”时，请求消息体中包含cluster_master_secret字段。 - 当“login_mode”配置为“1”时，请求消息体中包含node_public_cert_name字段。  枚举值： - 0：密码方式 - 1：密钥对方式
	LoginMode *CreateClusterReqV11LoginMode `json:"login_mode,omitempty"`

	// 节点列表信息。  说明：如下参数和该参数任选一组进行配置即可。  master_node_num、master_node_size、core_node_num、core_node_size、master_data_volume_type、master_data_volume_size、master_data_volume_count、core_data_volume_type、core_data_volume_size、core_data_volume_count、volume_type、volume_size、task_node_groups。
	NodeGroups *[]NodeGroupV11 `json:"node_groups,omitempty"`
}

func (o CreateClusterReqV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterReqV11 struct{}"
	}

	return strings.Join([]string{"CreateClusterReqV11", string(data)}, " ")
}

type CreateClusterReqV11BillingType struct {
	value int32
}

type CreateClusterReqV11BillingTypeEnum struct {
	E_12 CreateClusterReqV11BillingType
}

func GetCreateClusterReqV11BillingTypeEnum() CreateClusterReqV11BillingTypeEnum {
	return CreateClusterReqV11BillingTypeEnum{
		E_12: CreateClusterReqV11BillingType{
			value: 12,
		},
	}
}

func (c CreateClusterReqV11BillingType) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11BillingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11BillingType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateClusterReqV11VolumeType struct {
	value string
}

type CreateClusterReqV11VolumeTypeEnum struct {
	SATA  CreateClusterReqV11VolumeType
	SAS   CreateClusterReqV11VolumeType
	SSD   CreateClusterReqV11VolumeType
	GPSSD CreateClusterReqV11VolumeType
}

func GetCreateClusterReqV11VolumeTypeEnum() CreateClusterReqV11VolumeTypeEnum {
	return CreateClusterReqV11VolumeTypeEnum{
		SATA: CreateClusterReqV11VolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqV11VolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqV11VolumeType{
			value: "SSD",
		},
		GPSSD: CreateClusterReqV11VolumeType{
			value: "GPSSD",
		},
	}
}

func (c CreateClusterReqV11VolumeType) Value() string {
	return c.value
}

func (c CreateClusterReqV11VolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11VolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateClusterReqV11MasterDataVolumeType struct {
	value string
}

type CreateClusterReqV11MasterDataVolumeTypeEnum struct {
	SATA  CreateClusterReqV11MasterDataVolumeType
	SAS   CreateClusterReqV11MasterDataVolumeType
	SSD   CreateClusterReqV11MasterDataVolumeType
	GPSSD CreateClusterReqV11MasterDataVolumeType
}

func GetCreateClusterReqV11MasterDataVolumeTypeEnum() CreateClusterReqV11MasterDataVolumeTypeEnum {
	return CreateClusterReqV11MasterDataVolumeTypeEnum{
		SATA: CreateClusterReqV11MasterDataVolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqV11MasterDataVolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqV11MasterDataVolumeType{
			value: "SSD",
		},
		GPSSD: CreateClusterReqV11MasterDataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c CreateClusterReqV11MasterDataVolumeType) Value() string {
	return c.value
}

func (c CreateClusterReqV11MasterDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11MasterDataVolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateClusterReqV11MasterDataVolumeCount struct {
	value int32
}

type CreateClusterReqV11MasterDataVolumeCountEnum struct {
	E_1 CreateClusterReqV11MasterDataVolumeCount
}

func GetCreateClusterReqV11MasterDataVolumeCountEnum() CreateClusterReqV11MasterDataVolumeCountEnum {
	return CreateClusterReqV11MasterDataVolumeCountEnum{
		E_1: CreateClusterReqV11MasterDataVolumeCount{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11MasterDataVolumeCount) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11MasterDataVolumeCount) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11MasterDataVolumeCount) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateClusterReqV11CoreDataVolumeType struct {
	value string
}

type CreateClusterReqV11CoreDataVolumeTypeEnum struct {
	SATA  CreateClusterReqV11CoreDataVolumeType
	SAS   CreateClusterReqV11CoreDataVolumeType
	SSD   CreateClusterReqV11CoreDataVolumeType
	GPSSD CreateClusterReqV11CoreDataVolumeType
}

func GetCreateClusterReqV11CoreDataVolumeTypeEnum() CreateClusterReqV11CoreDataVolumeTypeEnum {
	return CreateClusterReqV11CoreDataVolumeTypeEnum{
		SATA: CreateClusterReqV11CoreDataVolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqV11CoreDataVolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqV11CoreDataVolumeType{
			value: "SSD",
		},
		GPSSD: CreateClusterReqV11CoreDataVolumeType{
			value: "GPSSD",
		},
	}
}

func (c CreateClusterReqV11CoreDataVolumeType) Value() string {
	return c.value
}

func (c CreateClusterReqV11CoreDataVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11CoreDataVolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateClusterReqV11SafeMode struct {
	value int32
}

type CreateClusterReqV11SafeModeEnum struct {
	E_0 CreateClusterReqV11SafeMode
	E_1 CreateClusterReqV11SafeMode
}

func GetCreateClusterReqV11SafeModeEnum() CreateClusterReqV11SafeModeEnum {
	return CreateClusterReqV11SafeModeEnum{
		E_0: CreateClusterReqV11SafeMode{
			value: 0,
		}, E_1: CreateClusterReqV11SafeMode{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11SafeMode) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11SafeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11SafeMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateClusterReqV11ClusterType struct {
	value int32
}

type CreateClusterReqV11ClusterTypeEnum struct {
	E_0 CreateClusterReqV11ClusterType
	E_1 CreateClusterReqV11ClusterType
}

func GetCreateClusterReqV11ClusterTypeEnum() CreateClusterReqV11ClusterTypeEnum {
	return CreateClusterReqV11ClusterTypeEnum{
		E_0: CreateClusterReqV11ClusterType{
			value: 0,
		}, E_1: CreateClusterReqV11ClusterType{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11ClusterType) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11ClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11ClusterType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateClusterReqV11LogCollection struct {
	value int32
}

type CreateClusterReqV11LogCollectionEnum struct {
	E_0 CreateClusterReqV11LogCollection
	E_1 CreateClusterReqV11LogCollection
}

func GetCreateClusterReqV11LogCollectionEnum() CreateClusterReqV11LogCollectionEnum {
	return CreateClusterReqV11LogCollectionEnum{
		E_0: CreateClusterReqV11LogCollection{
			value: 0,
		}, E_1: CreateClusterReqV11LogCollection{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11LogCollection) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11LogCollection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11LogCollection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type CreateClusterReqV11LoginMode struct {
	value int32
}

type CreateClusterReqV11LoginModeEnum struct {
	E_0 CreateClusterReqV11LoginMode
	E_1 CreateClusterReqV11LoginMode
}

func GetCreateClusterReqV11LoginModeEnum() CreateClusterReqV11LoginModeEnum {
	return CreateClusterReqV11LoginModeEnum{
		E_0: CreateClusterReqV11LoginMode{
			value: 0,
		}, E_1: CreateClusterReqV11LoginMode{
			value: 1,
		},
	}
}

func (c CreateClusterReqV11LoginMode) Value() int32 {
	return c.value
}

func (c CreateClusterReqV11LoginMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateClusterReqV11LoginMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
