package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDetailsOfInstanceV2Response Response Object
type ShowDetailsOfInstanceV2Response struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例所属租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态： - Creating：创建中 - CreateSuccess：创建成功 - CreateFail：创建失败 - Initing：初始化中 - Registering：注册中 - Running：运行中 - InitingFailed：初始化失败 - RegisterFailed：注册失败 - Installing：安装中 - InstallFailed：安装失败 - Updating：升级中 - UpdateFailed：升级失败 - Rollbacking：回滚中 - RollbackSuccess：回滚成功 - RollbackFailed：回滚失败 - Deleting：删除中 - DeleteFailed：删除失败 - Unregistering：注销中 - UnRegisterFailed：注销失败 - CreateTimeout：创建超时 - InitTimeout：初始化超时 - RegisterTimeout：注册超时 - InstallTimeout：安装超时 - UpdateTimeout：升级超时 - RollbackTimeout：回滚超时 - DeleteTimeout：删除超时 - UnregisterTimeout：注销超时 - Starting：启动中 - Freezing：冻结中 - Frozen：已冻结 - Restarting：重启中 - RestartFail：重启失败 - Unhealthy：实例异常 - RestartTimeout：重启超时 - Resizing：规格变更中 - ResizeFailed：规格变更失败 - ResizeTimeout：规格变更超时
	Status *ShowDetailsOfInstanceV2ResponseStatus `json:"status,omitempty"`

	// 实例状态对应编号 - 1：创建中 - 2：创建成功 - 3：创建失败 - 4：初始化中 - 5：注册中 - 6：运行中 - 7：初始化失败 - 8：注册失败 - 10：安装中 - 11：安装失败 - 12：升级中 - 13：升级失败 - 20：回滚中 - 21：回滚成功 - 22：回滚失败 - 23：删除中 - 24：删除失败 - 25：注销中 - 26：注销失败 - 27：创建超时 - 28：初始化超时 - 29：注册超时 - 30：安装超时 - 31：升级超时 - 32：回滚超时 - 33：删除超时 - 34：注销超时 - 35：启动中 - 36：冻结中 - 37：已冻结 - 38：重启中 - 39：重启失败 - 40：实例异常 - 41：重启超时 - 42：规格变更中 - 43：规格变更失败 - 44：规格变更超时
	InstanceStatus *ShowDetailsOfInstanceV2ResponseInstanceStatus `json:"instance_status,omitempty"`

	// 实例类型  默认apig
	Type *string `json:"type,omitempty"`

	// 实例规格： - BASIC：基础版实例 - PROFESSIONAL：专业版实例 - ENTERPRISE：企业版实例 - PLATINUM：铂金版实例 - BASIC_IPV6：基础版IPV6实例 - PROFESSIONAL_IPV6：专业版IPV6实例 - ENTERPRISE_IPV6：企业版IPV6实例 - PLATINUM_IPV6：铂金版IPV6实例
	Spec *ShowDetailsOfInstanceV2ResponseSpec `json:"spec,omitempty"`

	// 实例创建时间。unix时间戳格式。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 企业项目ID，企业账号必填
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例绑定的弹性IP地址
	EipAddress *string `json:"eip_address,omitempty"`

	// 实例计费方式： - 0：按需计费 - 1：[包周期计费](tag:hws)[暂未使用](tag:hws_hk,cmcc,ctc,DT,g42,hk_g42,hk_sbc,hk_tm,hws_eu,hws_ocb,OCB,sbc,tm)
	ChargingMode *ShowDetailsOfInstanceV2ResponseChargingMode `json:"charging_mode,omitempty"`

	// [包周期计费订单编号](tag:hws)[计费订单编号参数暂未使用](tag:hws_hk,cmcc,ctc,DT,g42,hk_g42,hk_sbc,hk_tm,hws_eu,hws_ocb,OCB,sbc,tm)
	CbcMetadata *string `json:"cbc_metadata,omitempty"`

	// 实例使用的负载均衡器类型 - lvs Linux虚拟服务器 - elb 弹性负载均衡，elb仅部分region支持
	LoadbalancerProvider *ShowDetailsOfInstanceV2ResponseLoadbalancerProvider `json:"loadbalancer_provider,omitempty"`

	// 云运营限制操作锁
	CbcOperationLocks *[]CbcOperationLock `json:"cbc_operation_locks,omitempty"`

	// 实例描述
	Description *string `json:"description,omitempty"`

	// 虚拟私有云ID。  获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询VPC列表”章节。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网的网络ID。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询子网列表”章节。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 指定实例所属的安全组。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询安全组列表”章节。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// '维护时间窗开始时间。时间格式为 xx:00:00，xx取值为02,06,10,14,18,22。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。'
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// '维护时间窗结束时间。时间格式为 xx:00:00，与维护时间窗开始时间相差4个小时。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次'。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 实例入口，虚拟私有云访问地址
	IngressIp *string `json:"ingress_ip,omitempty"`

	// 实例入口，虚拟私有云访问地址 (IPv6)
	IngressIpV6 *string `json:"ingress_ip_v6,omitempty"`

	// 实例所属用户ID
	UserId *string `json:"user_id,omitempty"`

	// 出公网网段 (IPv6)  。  当前仅部分region部分可用区支持IPv6
	NatEipIpv6Cidr *string `json:"nat_eip_ipv6_cidr,omitempty"`

	// 弹性IP地址(IPv6)。  当前仅部分region部分可用区支持IPv6
	EipIpv6Address *string `json:"eip_ipv6_address,omitempty"`

	// 实例出公网IP
	NatEipAddress *string `json:"nat_eip_address,omitempty"`

	// 出公网带宽
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 出公网带宽计费模式
	BandwidthChargingMode *string `json:"bandwidth_charging_mode,omitempty"`

	// 可用区
	AvailableZoneIds *string `json:"available_zone_ids,omitempty"`

	// 实例版本编号
	InstanceVersion *string `json:"instance_version,omitempty"`

	// 子网的网络ID。  暂不支持
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	// roma弹性公网IP。  暂不支持
	RomaEipAddress *string `json:"roma_eip_address,omitempty"`

	// 监听信息  暂不支持
	Listeners *interface{} `json:"listeners,omitempty"`

	// 实例支持的特性列表
	SupportedFeatures *[]string `json:"supported_features,omitempty"`

	EndpointService *EndpointService `json:"endpoint_service,omitempty"`

	// 终端节点服务列表
	EndpointServices *[]EndpointService `json:"endpoint_services,omitempty"`

	NodeIps *NodeIps `json:"node_ips,omitempty"`

	NodeIpv6Ips *NodeIps `json:"node_ipv6_ips,omitempty"`

	// 实例IPV6出公网IP
	PublicIpv6Ips *[]string `json:"public_ipv6_ips,omitempty"`

	// 公网入口地址列表
	Publicips *[]IpDetails `json:"publicips,omitempty"`

	// 私网入口地址列表
	Privateips *[]IpDetails `json:"privateips,omitempty"`

	// 不建议使用的实例入口列表：分组子域名解析未指向到该列表的IP，不能用该列表的IP直接访问DEFAULT分组
	UnreliableIps *[]string `json:"unreliable_ips,omitempty"`

	// 实例是否支持全量IPv6 - true：支持，实例入口和实例出口都支持IPV6 - false：不支持，仅实例入口支持IPV6
	EnableFullstackIpv6 *bool `json:"enable_fullstack_ipv6,omitempty"`

	// 实例是否可释放 - true：可释放 - false：不可释放
	IsReleasable *bool `json:"is_releasable,omitempty"`

	// 入公网带宽计费模式
	IngressBandwidthChargingMode *string `json:"ingress_bandwidth_charging_mode,omitempty"`
	HttpStatusCode               int     `json:"-"`
}

func (o ShowDetailsOfInstanceV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfInstanceV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfInstanceV2Response", string(data)}, " ")
}

type ShowDetailsOfInstanceV2ResponseStatus struct {
	value string
}

type ShowDetailsOfInstanceV2ResponseStatusEnum struct {
	CREATING           ShowDetailsOfInstanceV2ResponseStatus
	CREATE_SUCCESS     ShowDetailsOfInstanceV2ResponseStatus
	CREATE_FAIL        ShowDetailsOfInstanceV2ResponseStatus
	INITING            ShowDetailsOfInstanceV2ResponseStatus
	REGISTERING        ShowDetailsOfInstanceV2ResponseStatus
	RUNNING            ShowDetailsOfInstanceV2ResponseStatus
	INITING_FAILED     ShowDetailsOfInstanceV2ResponseStatus
	REGISTER_FAILED    ShowDetailsOfInstanceV2ResponseStatus
	INSTALLING         ShowDetailsOfInstanceV2ResponseStatus
	INSTALL_FAILED     ShowDetailsOfInstanceV2ResponseStatus
	UPDATING           ShowDetailsOfInstanceV2ResponseStatus
	UPDATE_FAILED      ShowDetailsOfInstanceV2ResponseStatus
	ROLLBACKING        ShowDetailsOfInstanceV2ResponseStatus
	ROLLBACK_SUCCESS   ShowDetailsOfInstanceV2ResponseStatus
	ROLLBACK_FAILED    ShowDetailsOfInstanceV2ResponseStatus
	DELETING           ShowDetailsOfInstanceV2ResponseStatus
	DELETE_FAILED      ShowDetailsOfInstanceV2ResponseStatus
	UNREGISTERING      ShowDetailsOfInstanceV2ResponseStatus
	UN_REGISTER_FAILED ShowDetailsOfInstanceV2ResponseStatus
	CREATE_TIMEOUT     ShowDetailsOfInstanceV2ResponseStatus
	INIT_TIMEOUT       ShowDetailsOfInstanceV2ResponseStatus
	REGISTER_TIMEOUT   ShowDetailsOfInstanceV2ResponseStatus
	INSTALL_TIMEOUT    ShowDetailsOfInstanceV2ResponseStatus
	UPDATE_TIMEOUT     ShowDetailsOfInstanceV2ResponseStatus
	ROLLBACK_TIMEOUT   ShowDetailsOfInstanceV2ResponseStatus
	DELETE_TIMEOUT     ShowDetailsOfInstanceV2ResponseStatus
	UNREGISTER_TIMEOUT ShowDetailsOfInstanceV2ResponseStatus
	STARTING           ShowDetailsOfInstanceV2ResponseStatus
	FREEZING           ShowDetailsOfInstanceV2ResponseStatus
	FROZEN             ShowDetailsOfInstanceV2ResponseStatus
	RESTARTING         ShowDetailsOfInstanceV2ResponseStatus
	RESTART_FAIL       ShowDetailsOfInstanceV2ResponseStatus
	UNHEALTHY          ShowDetailsOfInstanceV2ResponseStatus
	RESTART_TIMEOUT    ShowDetailsOfInstanceV2ResponseStatus
	RESIZING           ShowDetailsOfInstanceV2ResponseStatus
	RESIZE_FAILED      ShowDetailsOfInstanceV2ResponseStatus
	RESIZE_TIMEOUT     ShowDetailsOfInstanceV2ResponseStatus
}

func GetShowDetailsOfInstanceV2ResponseStatusEnum() ShowDetailsOfInstanceV2ResponseStatusEnum {
	return ShowDetailsOfInstanceV2ResponseStatusEnum{
		CREATING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Creating",
		},
		CREATE_SUCCESS: ShowDetailsOfInstanceV2ResponseStatus{
			value: "CreateSuccess",
		},
		CREATE_FAIL: ShowDetailsOfInstanceV2ResponseStatus{
			value: "CreateFail",
		},
		INITING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Initing",
		},
		REGISTERING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Registering",
		},
		RUNNING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Running",
		},
		INITING_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "InitingFailed",
		},
		REGISTER_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "RegisterFailed",
		},
		INSTALLING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Installing",
		},
		INSTALL_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "InstallFailed",
		},
		UPDATING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Updating",
		},
		UPDATE_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "UpdateFailed",
		},
		ROLLBACKING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Rollbacking",
		},
		ROLLBACK_SUCCESS: ShowDetailsOfInstanceV2ResponseStatus{
			value: "RollbackSuccess",
		},
		ROLLBACK_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "RollbackFailed",
		},
		DELETING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Deleting",
		},
		DELETE_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "DeleteFailed",
		},
		UNREGISTERING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Unregistering",
		},
		UN_REGISTER_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "UnRegisterFailed",
		},
		CREATE_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "CreateTimeout",
		},
		INIT_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "InitTimeout",
		},
		REGISTER_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "RegisterTimeout",
		},
		INSTALL_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "InstallTimeout",
		},
		UPDATE_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "UpdateTimeout",
		},
		ROLLBACK_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "RollbackTimeout",
		},
		DELETE_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "DeleteTimeout",
		},
		UNREGISTER_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "UnregisterTimeout",
		},
		STARTING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Starting",
		},
		FREEZING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Freezing",
		},
		FROZEN: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Frozen",
		},
		RESTARTING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Restarting",
		},
		RESTART_FAIL: ShowDetailsOfInstanceV2ResponseStatus{
			value: "RestartFail",
		},
		UNHEALTHY: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Unhealthy",
		},
		RESTART_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "RestartTimeout",
		},
		RESIZING: ShowDetailsOfInstanceV2ResponseStatus{
			value: "Resizing",
		},
		RESIZE_FAILED: ShowDetailsOfInstanceV2ResponseStatus{
			value: "ResizeFailed",
		},
		RESIZE_TIMEOUT: ShowDetailsOfInstanceV2ResponseStatus{
			value: "ResizeTimeout",
		},
	}
}

func (c ShowDetailsOfInstanceV2ResponseStatus) Value() string {
	return c.value
}

func (c ShowDetailsOfInstanceV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfInstanceV2ResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfInstanceV2ResponseInstanceStatus struct {
	value int32
}

type ShowDetailsOfInstanceV2ResponseInstanceStatusEnum struct {
	E_1  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_2  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_3  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_4  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_5  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_6  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_7  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_8  ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_10 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_11 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_12 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_13 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_20 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_21 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_22 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_23 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_24 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_25 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_26 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_27 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_28 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_29 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_30 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_31 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_32 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_33 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_34 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_35 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_36 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_37 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_38 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_39 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_40 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_41 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_42 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_43 ShowDetailsOfInstanceV2ResponseInstanceStatus
	E_44 ShowDetailsOfInstanceV2ResponseInstanceStatus
}

func GetShowDetailsOfInstanceV2ResponseInstanceStatusEnum() ShowDetailsOfInstanceV2ResponseInstanceStatusEnum {
	return ShowDetailsOfInstanceV2ResponseInstanceStatusEnum{
		E_1: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 1,
		}, E_2: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 2,
		}, E_3: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 3,
		}, E_4: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 4,
		}, E_5: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 5,
		}, E_6: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 6,
		}, E_7: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 7,
		}, E_8: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 8,
		}, E_10: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 10,
		}, E_11: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 11,
		}, E_12: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 12,
		}, E_13: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 13,
		}, E_20: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 20,
		}, E_21: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 21,
		}, E_22: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 22,
		}, E_23: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 23,
		}, E_24: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 24,
		}, E_25: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 25,
		}, E_26: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 26,
		}, E_27: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 27,
		}, E_28: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 28,
		}, E_29: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 29,
		}, E_30: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 30,
		}, E_31: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 31,
		}, E_32: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 32,
		}, E_33: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 33,
		}, E_34: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 34,
		}, E_35: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 35,
		}, E_36: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 36,
		}, E_37: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 37,
		}, E_38: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 38,
		}, E_39: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 39,
		}, E_40: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 40,
		}, E_41: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 41,
		}, E_42: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 42,
		}, E_43: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 43,
		}, E_44: ShowDetailsOfInstanceV2ResponseInstanceStatus{
			value: 44,
		},
	}
}

func (c ShowDetailsOfInstanceV2ResponseInstanceStatus) Value() int32 {
	return c.value
}

func (c ShowDetailsOfInstanceV2ResponseInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfInstanceV2ResponseInstanceStatus) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfInstanceV2ResponseSpec struct {
	value string
}

type ShowDetailsOfInstanceV2ResponseSpecEnum struct {
	BASIC             ShowDetailsOfInstanceV2ResponseSpec
	PROFESSIONAL      ShowDetailsOfInstanceV2ResponseSpec
	ENTERPRISE        ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM          ShowDetailsOfInstanceV2ResponseSpec
	BASIC_IPV6        ShowDetailsOfInstanceV2ResponseSpec
	PROFESSIONAL_IPV6 ShowDetailsOfInstanceV2ResponseSpec
	ENTERPRISE_IPV6   ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_IPV6     ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_X2       ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_X3       ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_X4       ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_X5       ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_X6       ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_X7       ShowDetailsOfInstanceV2ResponseSpec
	PLATINUM_X8       ShowDetailsOfInstanceV2ResponseSpec
}

func GetShowDetailsOfInstanceV2ResponseSpecEnum() ShowDetailsOfInstanceV2ResponseSpecEnum {
	return ShowDetailsOfInstanceV2ResponseSpecEnum{
		BASIC: ShowDetailsOfInstanceV2ResponseSpec{
			value: "BASIC",
		},
		PROFESSIONAL: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PROFESSIONAL",
		},
		ENTERPRISE: ShowDetailsOfInstanceV2ResponseSpec{
			value: "ENTERPRISE",
		},
		PLATINUM: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM",
		},
		BASIC_IPV6: ShowDetailsOfInstanceV2ResponseSpec{
			value: "BASIC_IPV6",
		},
		PROFESSIONAL_IPV6: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PROFESSIONAL_IPV6",
		},
		ENTERPRISE_IPV6: ShowDetailsOfInstanceV2ResponseSpec{
			value: "ENTERPRISE_IPV6",
		},
		PLATINUM_IPV6: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_IPV6",
		},
		PLATINUM_X2: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_X2",
		},
		PLATINUM_X3: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_X3",
		},
		PLATINUM_X4: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_X4",
		},
		PLATINUM_X5: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_X5",
		},
		PLATINUM_X6: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_X6",
		},
		PLATINUM_X7: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_X7",
		},
		PLATINUM_X8: ShowDetailsOfInstanceV2ResponseSpec{
			value: "PLATINUM_X8",
		},
	}
}

func (c ShowDetailsOfInstanceV2ResponseSpec) Value() string {
	return c.value
}

func (c ShowDetailsOfInstanceV2ResponseSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfInstanceV2ResponseSpec) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfInstanceV2ResponseChargingMode struct {
	value int32
}

type ShowDetailsOfInstanceV2ResponseChargingModeEnum struct {
	E_0 ShowDetailsOfInstanceV2ResponseChargingMode
	E_1 ShowDetailsOfInstanceV2ResponseChargingMode
}

func GetShowDetailsOfInstanceV2ResponseChargingModeEnum() ShowDetailsOfInstanceV2ResponseChargingModeEnum {
	return ShowDetailsOfInstanceV2ResponseChargingModeEnum{
		E_0: ShowDetailsOfInstanceV2ResponseChargingMode{
			value: 0,
		}, E_1: ShowDetailsOfInstanceV2ResponseChargingMode{
			value: 1,
		},
	}
}

func (c ShowDetailsOfInstanceV2ResponseChargingMode) Value() int32 {
	return c.value
}

func (c ShowDetailsOfInstanceV2ResponseChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfInstanceV2ResponseChargingMode) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfInstanceV2ResponseLoadbalancerProvider struct {
	value string
}

type ShowDetailsOfInstanceV2ResponseLoadbalancerProviderEnum struct {
	LVS ShowDetailsOfInstanceV2ResponseLoadbalancerProvider
	ELB ShowDetailsOfInstanceV2ResponseLoadbalancerProvider
}

func GetShowDetailsOfInstanceV2ResponseLoadbalancerProviderEnum() ShowDetailsOfInstanceV2ResponseLoadbalancerProviderEnum {
	return ShowDetailsOfInstanceV2ResponseLoadbalancerProviderEnum{
		LVS: ShowDetailsOfInstanceV2ResponseLoadbalancerProvider{
			value: "lvs",
		},
		ELB: ShowDetailsOfInstanceV2ResponseLoadbalancerProvider{
			value: "elb",
		},
	}
}

func (c ShowDetailsOfInstanceV2ResponseLoadbalancerProvider) Value() string {
	return c.value
}

func (c ShowDetailsOfInstanceV2ResponseLoadbalancerProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfInstanceV2ResponseLoadbalancerProvider) UnmarshalJSON(b []byte) error {
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
