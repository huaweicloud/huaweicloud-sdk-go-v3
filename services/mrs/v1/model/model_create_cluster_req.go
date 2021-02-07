package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type CreateClusterReq struct {
	// 集群版本。 MRS目前支持MRS 1.8.10、MRS 2.0.5、MRS 2.1.0版本。默认值为当前最新版本。
	ClusterVersion *string `json:"cluster_version,omitempty"`
	// 集群名称，不允许相同。 只能由字母、数字、中划线和下划线组成，并且长度为1～64个字符。
	ClusterName string `json:"cluster_name"`
	// Master节点数量。启用集群高可用功能时配置为2，不启用集群高可用功能时配置为1。
	MasterNodeNum CreateClusterReqMasterNodeNum `json:"master_node_num"`
	// Core节点数量。  取值范围：[1～500]  Core节点默认的最大值为500，如果用户需要的Core节点数大于500，请申请扩大配额，具体请参考[关于配额](https://support.huaweicloud.com/usermanual-iaas/zh-cn_topic_0040259342.html)
	CoreNodeNum int32 `json:"core_node_num"`
	// 集群的计费模式。  11：表示包年/包月。 12：表示按需计费。接口调用仅支持创建按需计费集群
	BillingType CreateClusterReqBillingType `json:"billing_type"`
	// 集群区域信息，请参见[终端节点及区域](https://developer.huaweicloud.com/endpoint?MRS)。
	DataCenter string `json:"data_center"`
	// 子网所在VPC名称。 通过VPC管理控制台获取名称： 登录管理控制台。 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 在“虚拟私有云”页面的列表中即可获取VPC名称。
	Vpc string `json:"vpc"`
	// Master节点的实例规格，例如：c3.4xlarge.2.linux.bigdata。MRS当前支持主机规格的配型由CPU+内存+Disk共同决定。实例规格详细说明请参见[MRS所使用的弹性云服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9006.html)和[MRS所使用的裸金属服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9001.html)。
	MasterNodeSize string `json:"master_node_size"`
	// Core节点的实例规格，例如：c3.4xlarge.2.linux.bigdata。实例规格详细说明请参见[MRS所使用的弹性云服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9006.html)和[MRS所使用的裸金属服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9001.html)。
	CoreNodeSize string `json:"core_node_size"`
	// 服务组件安装列表信息。参数说明请参见[表5](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__te1288dba79844d3fa5973939a3739d34)。
	ComponentList []ComponentList `json:"component_list"`
	// 可用分区ID。  华北-北京一可用区1（cn-north-1a）：ae04cf9d61544df3806a3feeb401b204 华北-北京一可用区2（cn-north-1b）：d573142f24894ef3bd3664de068b44b0 华东-上海二可用区1（cn-east-2a）：72d50cedc49846b9b42c21495f38d81c 华东-上海二可用区2（cn-east-2b）：38b0f7a602344246bcb0da47b5d548e7 华东-上海二可用区3（cn-east-2c）：5547fd6bf8f84bb5a7f9db062ad3d015 华南-广州可用区1（cn-south-1a）：34f5ff4865cf4ed6b270f15382ebdec5 华南-广州可用区2（cn-south-2b）：043c7e39ecb347a08dc8fcb6c35a274e 华南-广州可用区3（cn-south-1c）：af1687643e8c4ec1b34b688e4e3b8901 华北-北京四可用区1（cn-north-4a）：effdcbc7d4d64a02aa1fa26b42f56533 华北-北京四可用区2（cn-north-4b）：a0865121f83b41cbafce65930a22a6e8 华北-北京四可用区3（cn-north-4c）：2dcb154ac2724a6d92e9bcc859657c1e
	AvailableZoneId string `json:"available_zone_id"`
	// 子网所在VPC ID。 通过VPC管理控制台获取ID： 登录管理控制台。 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 在“虚拟私有云”页面的列表中即可获取VPC ID。
	VpcId string `json:"vpc_id"`
	// 网络ID。 通过VPC管理控制台获取网络ID： 登录管理控制台。 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 在“虚拟私有云”页面的列表中即可获取VPC网络ID。
	SubnetId string `json:"subnet_id"`
	// 子网名称。 通过VPC管理控制台获取子网名称： 登录管理控制台。 单击“虚拟私有云”，从左侧列表选择虚拟私有云。 在“虚拟私有云”页面的列表中即可获取VPC子网名称。
	SubnetName string `json:"subnet_name"`
	// 集群安全组的ID。 当该ID为空时MRS后台会自己创建安全组，自动创建的安全组名称以mrs_{cluster_name}开头。 当该ID不为空时，表示使用固定安全组来创建集群，传入的ID必须是当前租户中包含的安全组ID，且该安全组中包含一条全部协议，全部端口，源地址为指定的管理面节点IP的入方向规则。
	SecurityGroupsId *string `json:"security_groups_id,omitempty"`
	// 创建集群时可同时提交作业，当前版本暂时只支持新增一个作业，作业参数请参见[表6](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__t8ded0b3ae11742cea98a467ce26fd093)。
	AddJobs *[]AddJobs `json:"add_jobs,omitempty"`
	// Master和Core节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。可以根据如下应用场景合理选择磁盘存储空间大小： 数据存储和计算分离，数据存储在OBS系统中，集群费用相对较低，计算性能不高，并且集群随时可以删除，建议数据计算不频繁场景下使用。 数据存储和计算不分离，数据存储在HDFS中，集群费用相对较高，计算性能高，集群需要长期存在，建议数据计算频繁场景下使用。 取值范围：100GB～32000GB,传值只需填数字,不需要带单位GB 不建议使用该参数，详情请参考volume_type参数的说明。
	VolumeSize *int32 `json:"volume_size,omitempty"`
	// Master和Core节点的磁盘存储类别，目前支持SATA、SAS和SSD。磁盘参数可以使用volume_type和volume_size表示，也可以使用多磁盘相关的参数表示。volume_type和volume_size这两个参数如果与多磁盘参数同时出现，系统优先读取volume_type和volume_size参数。建议使用多磁盘参数。 SATA：普通IO SAS：高IO SSD：超高IO
	VolumeType *CreateClusterReqVolumeType `json:"volume_type,omitempty"`
	// 该参数为多磁盘参数，表示Master节点数据磁盘存储类别，目前支持SATA、SAS和SSD。
	MasterDataVolumeType *CreateClusterReqMasterDataVolumeType `json:"master_data_volume_type,omitempty"`
	// 该参数为多磁盘参数，表示Master节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。 取值范围：100GB～32000GB,传值只需填数字,不需要带单位GB
	MasterDataVolumeSize *int32 `json:"master_data_volume_size,omitempty"`
	// 该参数为多磁盘参数，表示Master节点数据磁盘个数。 取值只能是1
	MasterDataVolumeCount *CreateClusterReqMasterDataVolumeCount `json:"master_data_volume_count,omitempty"`
	// 该参数为多磁盘参数，表示Core节点数据磁盘存储类别，目前支持SATA、SAS和SSD。
	CoreDataVolumeType *CreateClusterReqCoreDataVolumeType `json:"core_data_volume_type,omitempty"`
	// 该参数为多磁盘参数，表示Core节点数据磁盘存储空间。为增大数据存储容量，创建集群时可同时添加磁盘。 取值范围：100GB～32000GB,传值只需填数字,不需要带单位GB
	CoreDataVolumeSize *int32 `json:"core_data_volume_size,omitempty"`
	// 该参数为多磁盘参数，表示Core节点数据磁盘个数。 取值范围：1～10
	CoreDataVolumeCount *int32 `json:"core_data_volume_count,omitempty"`
	// Task节点列表信息。参数说明请参见[表4](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__tc6bfa2a3d7a348d786a901f3a9327b50)。
	TaskNodeGroups *[]TaskNodeGroups `json:"task_node_groups,omitempty"`
	// 配置引导操作脚本信息。参数说明请参见[表13](https://support.huaweicloud.com/api-mrs/mrs_02_0028.html#mrs_02_0028__table1258382865010)。 MRS 1.7.1及以后版本支持该参数。
	BootstrapScripts *[]BootstrapScript `json:"bootstrap_scripts,omitempty"`
	// 密钥对名称。用户可以使用密钥对方式登录集群节点。
	NodePublicCertName *string `json:"node_public_cert_name,omitempty"`
	// 配置MRS Manager管理员用户的密码。 密码长度应在8～32个字符之间 必须包含如下5种中至少3种字符的组合 至少一个小写字母 至少一个大写字母 至少一个数字 至少一个特殊字符：`~!@#$%^&*()-_=+\\|[{}];:'\",<.>/? 空格 不能与用户名或者倒序用户名相同 说明： 针对MRS 1.8.0以前的版本，仅当“safe_mode”配置为“1”时需要配置此参数。 针对MRS 1.8.0及以后版本，该参数为必选参数，不受参数“safe_mode”配置的影响。
	ClusterAdminSecret *string `json:"cluster_admin_secret,omitempty"`
	// 配置访问集群节点的root密码。 密码设置约束如下： 字符串类型，可输入的字符串长度为8-26。 至少包含三种字符组合，如大写字母，小写字母，数字，特殊字符（!@$%^-_=+[{}]:,./?），但不能包含空格。 不能与用户名或者倒序用户名相同。
	ClusterMasterSecret *string `json:"cluster_master_secret,omitempty"`
	// MRS集群运行模式。 0：普通集群，表示Kerberos认证关闭，用户可使用集群提供的所有功能。 1：安全集群，表示Kerberos认证开启，普通用户无权限使用MRS集群的“文件管理”和“作业管理”功能，并且无法查看Hadoop、Spark的作业记录以及集群资源使用情况。如果需要使用集群更多功能，需要找MRS Manager的管理员分配权限。 说明： 针对MRS 1.8.0以前的版本，仅当“safe_mode”配置为“1”时，请求消息体中包含cluster_admin_secret字段。 针对MRS 1.8.0及以后版本，请求消息体中包含cluster_admin_secret字段，不受参数“safe_mode”配置的影响。
	SafeMode CreateClusterReqSafeMode `json:"safe_mode"`
	// 集群类型。 0：分析集群 1：流式集群 默认值为0：分析集群。 说明：暂不支持通过接口方式创建混合集群。
	ClusterType *CreateClusterReqClusterType `json:"cluster_type,omitempty"`
	// 集群创建失败时，是否收集失败日志。  0：不收集 1：收集 默认设置为1，将创建OBS桶仅用于MRS集群创建失败时的日志收集。
	LogCollection *CreateClusterReqLogCollection `json:"log_collection,omitempty"`
	// 集群的标签信息。 同一个集群最多能使用10个tag，tag的名称（key）不能重复 标签的键/值不能包含“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”
	Tags *[]Tag `json:"tags,omitempty"`
	// 集群登录方式。 0：密码方式 1：密钥对方式 默认设置为1。 当“login_mode”配置为“0”时，请求消息体中包含cluster_master_secret字段。 当“login_mode”配置为“1”时，请求消息体中包含node_public_cert_name字段。 说明： 该参数仅适用于MRS 1.6.2及以后版本的集群，MRS 1.6.2前的版本不支持。
	LoginMode *CreateClusterReqLoginMode `json:"login_mode,omitempty"`
}

func (o CreateClusterReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateClusterReq struct{}"
	}

	return strings.Join([]string{"CreateClusterReq", string(data)}, " ")
}

type CreateClusterReqMasterNodeNum struct {
	value int32
}

type CreateClusterReqMasterNodeNumEnum struct {
	E_1 CreateClusterReqMasterNodeNum
	E_2 CreateClusterReqMasterNodeNum
}

func GetCreateClusterReqMasterNodeNumEnum() CreateClusterReqMasterNodeNumEnum {
	return CreateClusterReqMasterNodeNumEnum{
		E_1: CreateClusterReqMasterNodeNum{
			value: 1,
		}, E_2: CreateClusterReqMasterNodeNum{
			value: 2,
		},
	}
}

func (c CreateClusterReqMasterNodeNum) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqMasterNodeNum) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqBillingType struct {
	value int32
}

type CreateClusterReqBillingTypeEnum struct {
	E_12 CreateClusterReqBillingType
}

func GetCreateClusterReqBillingTypeEnum() CreateClusterReqBillingTypeEnum {
	return CreateClusterReqBillingTypeEnum{
		E_12: CreateClusterReqBillingType{
			value: 12,
		},
	}
}

func (c CreateClusterReqBillingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqBillingType) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqVolumeType struct {
	value string
}

type CreateClusterReqVolumeTypeEnum struct {
	SATA CreateClusterReqVolumeType
	SAS  CreateClusterReqVolumeType
	SSD  CreateClusterReqVolumeType
}

func GetCreateClusterReqVolumeTypeEnum() CreateClusterReqVolumeTypeEnum {
	return CreateClusterReqVolumeTypeEnum{
		SATA: CreateClusterReqVolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqVolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqVolumeType{
			value: "SSD",
		},
	}
}

func (c CreateClusterReqVolumeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqVolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateClusterReqMasterDataVolumeType struct {
	value string
}

type CreateClusterReqMasterDataVolumeTypeEnum struct {
	SATA CreateClusterReqMasterDataVolumeType
	SAS  CreateClusterReqMasterDataVolumeType
	SSD  CreateClusterReqMasterDataVolumeType
}

func GetCreateClusterReqMasterDataVolumeTypeEnum() CreateClusterReqMasterDataVolumeTypeEnum {
	return CreateClusterReqMasterDataVolumeTypeEnum{
		SATA: CreateClusterReqMasterDataVolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqMasterDataVolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqMasterDataVolumeType{
			value: "SSD",
		},
	}
}

func (c CreateClusterReqMasterDataVolumeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqMasterDataVolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateClusterReqMasterDataVolumeCount struct {
	value int32
}

type CreateClusterReqMasterDataVolumeCountEnum struct {
	E_1 CreateClusterReqMasterDataVolumeCount
}

func GetCreateClusterReqMasterDataVolumeCountEnum() CreateClusterReqMasterDataVolumeCountEnum {
	return CreateClusterReqMasterDataVolumeCountEnum{
		E_1: CreateClusterReqMasterDataVolumeCount{
			value: 1,
		},
	}
}

func (c CreateClusterReqMasterDataVolumeCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqMasterDataVolumeCount) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqCoreDataVolumeType struct {
	value string
}

type CreateClusterReqCoreDataVolumeTypeEnum struct {
	SATA CreateClusterReqCoreDataVolumeType
	SAS  CreateClusterReqCoreDataVolumeType
	SSD  CreateClusterReqCoreDataVolumeType
}

func GetCreateClusterReqCoreDataVolumeTypeEnum() CreateClusterReqCoreDataVolumeTypeEnum {
	return CreateClusterReqCoreDataVolumeTypeEnum{
		SATA: CreateClusterReqCoreDataVolumeType{
			value: "SATA",
		},
		SAS: CreateClusterReqCoreDataVolumeType{
			value: "SAS",
		},
		SSD: CreateClusterReqCoreDataVolumeType{
			value: "SSD",
		},
	}
}

func (c CreateClusterReqCoreDataVolumeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqCoreDataVolumeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateClusterReqSafeMode struct {
	value int32
}

type CreateClusterReqSafeModeEnum struct {
	E_0 CreateClusterReqSafeMode
	E_1 CreateClusterReqSafeMode
}

func GetCreateClusterReqSafeModeEnum() CreateClusterReqSafeModeEnum {
	return CreateClusterReqSafeModeEnum{
		E_0: CreateClusterReqSafeMode{
			value: 0,
		}, E_1: CreateClusterReqSafeMode{
			value: 1,
		},
	}
}

func (c CreateClusterReqSafeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqSafeMode) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqClusterType struct {
	value int32
}

type CreateClusterReqClusterTypeEnum struct {
	E_0 CreateClusterReqClusterType
	E_1 CreateClusterReqClusterType
}

func GetCreateClusterReqClusterTypeEnum() CreateClusterReqClusterTypeEnum {
	return CreateClusterReqClusterTypeEnum{
		E_0: CreateClusterReqClusterType{
			value: 0,
		}, E_1: CreateClusterReqClusterType{
			value: 1,
		},
	}
}

func (c CreateClusterReqClusterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqClusterType) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqLogCollection struct {
	value int32
}

type CreateClusterReqLogCollectionEnum struct {
	E_0 CreateClusterReqLogCollection
	E_1 CreateClusterReqLogCollection
}

func GetCreateClusterReqLogCollectionEnum() CreateClusterReqLogCollectionEnum {
	return CreateClusterReqLogCollectionEnum{
		E_0: CreateClusterReqLogCollection{
			value: 0,
		}, E_1: CreateClusterReqLogCollection{
			value: 1,
		},
	}
}

func (c CreateClusterReqLogCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqLogCollection) UnmarshalJSON(b []byte) error {
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

type CreateClusterReqLoginMode struct {
	value int32
}

type CreateClusterReqLoginModeEnum struct {
	E_0 CreateClusterReqLoginMode
	E_1 CreateClusterReqLoginMode
}

func GetCreateClusterReqLoginModeEnum() CreateClusterReqLoginModeEnum {
	return CreateClusterReqLoginModeEnum{
		E_0: CreateClusterReqLoginMode{
			value: 0,
		}, E_1: CreateClusterReqLoginMode{
			value: 1,
		},
	}
}

func (c CreateClusterReqLoginMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateClusterReqLoginMode) UnmarshalJSON(b []byte) error {
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
