package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RespInstanceBase struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例所属租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态： - Creating：创建中 - CreateSuccess：创建成功 - CreateFail：创建失败 - Initing：初始化中 - Registering：注册中 - Running：运行中 - InitingFailed：初始化失败 - RegisterFailed：注册失败 - Installing：安装中 - InstallFailed：安装失败 - Updating：升级中 - UpdateFailed：升级失败 - Rollbacking：回滚中 - RollbackSuccess：回滚成功 - RollbackFailed：回滚失败 - Deleting：删除中 - DeleteFailed：删除失败 - Unregistering：注销中 - UnRegisterFailed：注销失败 - CreateTimeout：创建超时 - InitTimeout：初始化超时 - RegisterTimeout：注册超时 - InstallTimeout：安装超时 - UpdateTimeout：升级超时 - RollbackTimeout：回滚超时 - DeleteTimeout：删除超时 - UnregisterTimeout：注销超时 - Starting：启动中 - Freezing：冻结中 - Frozen：已冻结 - Restarting：重启中 - RestartFail：重启失败 - Unhealthy：实例异常 - RestartTimeout：重启超时 - Resizing：规格变更中 - ResizeFailed：规格变更失败 - ResizeTimeout：规格变更超时
	Status *RespInstanceBaseStatus `json:"status,omitempty"`

	// 实例状态对应编号 - 1：创建中 - 2：创建成功 - 3：创建失败 - 4：初始化中 - 5：注册中 - 6：运行中 - 7：初始化失败 - 8：注册失败 - 10：安装中 - 11：安装失败 - 12：升级中 - 13：升级失败 - 20：回滚中 - 21：回滚成功 - 22：回滚失败 - 23：删除中 - 24：删除失败 - 25：注销中 - 26：注销失败 - 27：创建超时 - 28：初始化超时 - 29：注册超时 - 30：安装超时 - 31：升级超时 - 32：回滚超时 - 33：删除超时 - 34：注销超时 - 35：启动中 - 36：冻结中 - 37：已冻结 - 38：重启中 - 39：重启失败 - 40：实例异常 - 41：重启超时 - 42：规格变更中 - 43：规格变更失败 - 44：规格变更超时
	InstanceStatus *RespInstanceBaseInstanceStatus `json:"instance_status,omitempty"`

	// 实例类型  默认apig
	Type *string `json:"type,omitempty"`

	// 实例规格： - BASIC：基础版实例 - PROFESSIONAL：专业版实例 - ENTERPRISE：企业版实例 - PLATINUM：铂金版实例 - BASIC_IPV6：基础版IPV6实例 - PROFESSIONAL_IPV6：专业版IPV6实例 - ENTERPRISE_IPV6：企业版IPV6实例 - PLATINUM_IPV6：铂金版IPV6实例
	Spec *RespInstanceBaseSpec `json:"spec,omitempty"`

	// 实例创建时间。unix时间戳格式。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 企业项目ID，企业帐号必填
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例绑定的弹性IP地址
	EipAddress *string `json:"eip_address,omitempty"`

	// 实例计费方式： - 0：按需计费 - 1：[包周期计费](tag:hws)[暂未使用](tag:hws_hk,cmcc,ctc,DT,g42,hk_g42,hk_sbc,hk_tm,hws_eu,hws_ocb,OCB,sbc,tm)
	ChargingMode *RespInstanceBaseChargingMode `json:"charging_mode,omitempty"`

	// [包周期计费订单编号](tag:hws)[计费订单编号参数暂未使用](tag:hws_hk,cmcc,ctc,DT,g42,hk_g42,hk_sbc,hk_tm,hws_eu,hws_ocb,OCB,sbc,tm)
	CbcMetadata *string `json:"cbc_metadata,omitempty"`

	// 实例使用的负载均衡器类型 - lvs Linux虚拟服务器 - elb 弹性负载均衡，elb仅部分region支持
	LoadbalancerProvider *RespInstanceBaseLoadbalancerProvider `json:"loadbalancer_provider,omitempty"`

	// 云运营限制操作锁
	CbcOperationLocks *[]CbcOperationLock `json:"cbc_operation_locks,omitempty"`
}

func (o RespInstanceBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespInstanceBase struct{}"
	}

	return strings.Join([]string{"RespInstanceBase", string(data)}, " ")
}

type RespInstanceBaseStatus struct {
	value string
}

type RespInstanceBaseStatusEnum struct {
	CREATING           RespInstanceBaseStatus
	CREATE_SUCCESS     RespInstanceBaseStatus
	CREATE_FAIL        RespInstanceBaseStatus
	INITING            RespInstanceBaseStatus
	REGISTERING        RespInstanceBaseStatus
	RUNNING            RespInstanceBaseStatus
	INITING_FAILED     RespInstanceBaseStatus
	REGISTER_FAILED    RespInstanceBaseStatus
	INSTALLING         RespInstanceBaseStatus
	INSTALL_FAILED     RespInstanceBaseStatus
	UPDATING           RespInstanceBaseStatus
	UPDATE_FAILED      RespInstanceBaseStatus
	ROLLBACKING        RespInstanceBaseStatus
	ROLLBACK_SUCCESS   RespInstanceBaseStatus
	ROLLBACK_FAILED    RespInstanceBaseStatus
	DELETING           RespInstanceBaseStatus
	DELETE_FAILED      RespInstanceBaseStatus
	UNREGISTERING      RespInstanceBaseStatus
	UN_REGISTER_FAILED RespInstanceBaseStatus
	CREATE_TIMEOUT     RespInstanceBaseStatus
	INIT_TIMEOUT       RespInstanceBaseStatus
	REGISTER_TIMEOUT   RespInstanceBaseStatus
	INSTALL_TIMEOUT    RespInstanceBaseStatus
	UPDATE_TIMEOUT     RespInstanceBaseStatus
	ROLLBACK_TIMEOUT   RespInstanceBaseStatus
	DELETE_TIMEOUT     RespInstanceBaseStatus
	UNREGISTER_TIMEOUT RespInstanceBaseStatus
	STARTING           RespInstanceBaseStatus
	FREEZING           RespInstanceBaseStatus
	FROZEN             RespInstanceBaseStatus
	RESTARTING         RespInstanceBaseStatus
	RESTART_FAIL       RespInstanceBaseStatus
	UNHEALTHY          RespInstanceBaseStatus
	RESTART_TIMEOUT    RespInstanceBaseStatus
	RESIZING           RespInstanceBaseStatus
	RESIZE_FAILED      RespInstanceBaseStatus
	RESIZE_TIMEOUT     RespInstanceBaseStatus
}

func GetRespInstanceBaseStatusEnum() RespInstanceBaseStatusEnum {
	return RespInstanceBaseStatusEnum{
		CREATING: RespInstanceBaseStatus{
			value: "Creating",
		},
		CREATE_SUCCESS: RespInstanceBaseStatus{
			value: "CreateSuccess",
		},
		CREATE_FAIL: RespInstanceBaseStatus{
			value: "CreateFail",
		},
		INITING: RespInstanceBaseStatus{
			value: "Initing",
		},
		REGISTERING: RespInstanceBaseStatus{
			value: "Registering",
		},
		RUNNING: RespInstanceBaseStatus{
			value: "Running",
		},
		INITING_FAILED: RespInstanceBaseStatus{
			value: "InitingFailed",
		},
		REGISTER_FAILED: RespInstanceBaseStatus{
			value: "RegisterFailed",
		},
		INSTALLING: RespInstanceBaseStatus{
			value: "Installing",
		},
		INSTALL_FAILED: RespInstanceBaseStatus{
			value: "InstallFailed",
		},
		UPDATING: RespInstanceBaseStatus{
			value: "Updating",
		},
		UPDATE_FAILED: RespInstanceBaseStatus{
			value: "UpdateFailed",
		},
		ROLLBACKING: RespInstanceBaseStatus{
			value: "Rollbacking",
		},
		ROLLBACK_SUCCESS: RespInstanceBaseStatus{
			value: "RollbackSuccess",
		},
		ROLLBACK_FAILED: RespInstanceBaseStatus{
			value: "RollbackFailed",
		},
		DELETING: RespInstanceBaseStatus{
			value: "Deleting",
		},
		DELETE_FAILED: RespInstanceBaseStatus{
			value: "DeleteFailed",
		},
		UNREGISTERING: RespInstanceBaseStatus{
			value: "Unregistering",
		},
		UN_REGISTER_FAILED: RespInstanceBaseStatus{
			value: "UnRegisterFailed",
		},
		CREATE_TIMEOUT: RespInstanceBaseStatus{
			value: "CreateTimeout",
		},
		INIT_TIMEOUT: RespInstanceBaseStatus{
			value: "InitTimeout",
		},
		REGISTER_TIMEOUT: RespInstanceBaseStatus{
			value: "RegisterTimeout",
		},
		INSTALL_TIMEOUT: RespInstanceBaseStatus{
			value: "InstallTimeout",
		},
		UPDATE_TIMEOUT: RespInstanceBaseStatus{
			value: "UpdateTimeout",
		},
		ROLLBACK_TIMEOUT: RespInstanceBaseStatus{
			value: "RollbackTimeout",
		},
		DELETE_TIMEOUT: RespInstanceBaseStatus{
			value: "DeleteTimeout",
		},
		UNREGISTER_TIMEOUT: RespInstanceBaseStatus{
			value: "UnregisterTimeout",
		},
		STARTING: RespInstanceBaseStatus{
			value: "Starting",
		},
		FREEZING: RespInstanceBaseStatus{
			value: "Freezing",
		},
		FROZEN: RespInstanceBaseStatus{
			value: "Frozen",
		},
		RESTARTING: RespInstanceBaseStatus{
			value: "Restarting",
		},
		RESTART_FAIL: RespInstanceBaseStatus{
			value: "RestartFail",
		},
		UNHEALTHY: RespInstanceBaseStatus{
			value: "Unhealthy",
		},
		RESTART_TIMEOUT: RespInstanceBaseStatus{
			value: "RestartTimeout",
		},
		RESIZING: RespInstanceBaseStatus{
			value: "Resizing",
		},
		RESIZE_FAILED: RespInstanceBaseStatus{
			value: "ResizeFailed",
		},
		RESIZE_TIMEOUT: RespInstanceBaseStatus{
			value: "ResizeTimeout",
		},
	}
}

func (c RespInstanceBaseStatus) Value() string {
	return c.value
}

func (c RespInstanceBaseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespInstanceBaseStatus) UnmarshalJSON(b []byte) error {
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

type RespInstanceBaseInstanceStatus struct {
	value int32
}

type RespInstanceBaseInstanceStatusEnum struct {
	E_1  RespInstanceBaseInstanceStatus
	E_2  RespInstanceBaseInstanceStatus
	E_3  RespInstanceBaseInstanceStatus
	E_4  RespInstanceBaseInstanceStatus
	E_5  RespInstanceBaseInstanceStatus
	E_6  RespInstanceBaseInstanceStatus
	E_7  RespInstanceBaseInstanceStatus
	E_8  RespInstanceBaseInstanceStatus
	E_10 RespInstanceBaseInstanceStatus
	E_11 RespInstanceBaseInstanceStatus
	E_12 RespInstanceBaseInstanceStatus
	E_13 RespInstanceBaseInstanceStatus
	E_20 RespInstanceBaseInstanceStatus
	E_21 RespInstanceBaseInstanceStatus
	E_22 RespInstanceBaseInstanceStatus
	E_23 RespInstanceBaseInstanceStatus
	E_24 RespInstanceBaseInstanceStatus
	E_25 RespInstanceBaseInstanceStatus
	E_26 RespInstanceBaseInstanceStatus
	E_27 RespInstanceBaseInstanceStatus
	E_28 RespInstanceBaseInstanceStatus
	E_29 RespInstanceBaseInstanceStatus
	E_30 RespInstanceBaseInstanceStatus
	E_31 RespInstanceBaseInstanceStatus
	E_32 RespInstanceBaseInstanceStatus
	E_33 RespInstanceBaseInstanceStatus
	E_34 RespInstanceBaseInstanceStatus
	E_35 RespInstanceBaseInstanceStatus
	E_36 RespInstanceBaseInstanceStatus
	E_37 RespInstanceBaseInstanceStatus
	E_38 RespInstanceBaseInstanceStatus
	E_39 RespInstanceBaseInstanceStatus
	E_40 RespInstanceBaseInstanceStatus
	E_41 RespInstanceBaseInstanceStatus
	E_42 RespInstanceBaseInstanceStatus
	E_43 RespInstanceBaseInstanceStatus
	E_44 RespInstanceBaseInstanceStatus
}

func GetRespInstanceBaseInstanceStatusEnum() RespInstanceBaseInstanceStatusEnum {
	return RespInstanceBaseInstanceStatusEnum{
		E_1: RespInstanceBaseInstanceStatus{
			value: 1,
		}, E_2: RespInstanceBaseInstanceStatus{
			value: 2,
		}, E_3: RespInstanceBaseInstanceStatus{
			value: 3,
		}, E_4: RespInstanceBaseInstanceStatus{
			value: 4,
		}, E_5: RespInstanceBaseInstanceStatus{
			value: 5,
		}, E_6: RespInstanceBaseInstanceStatus{
			value: 6,
		}, E_7: RespInstanceBaseInstanceStatus{
			value: 7,
		}, E_8: RespInstanceBaseInstanceStatus{
			value: 8,
		}, E_10: RespInstanceBaseInstanceStatus{
			value: 10,
		}, E_11: RespInstanceBaseInstanceStatus{
			value: 11,
		}, E_12: RespInstanceBaseInstanceStatus{
			value: 12,
		}, E_13: RespInstanceBaseInstanceStatus{
			value: 13,
		}, E_20: RespInstanceBaseInstanceStatus{
			value: 20,
		}, E_21: RespInstanceBaseInstanceStatus{
			value: 21,
		}, E_22: RespInstanceBaseInstanceStatus{
			value: 22,
		}, E_23: RespInstanceBaseInstanceStatus{
			value: 23,
		}, E_24: RespInstanceBaseInstanceStatus{
			value: 24,
		}, E_25: RespInstanceBaseInstanceStatus{
			value: 25,
		}, E_26: RespInstanceBaseInstanceStatus{
			value: 26,
		}, E_27: RespInstanceBaseInstanceStatus{
			value: 27,
		}, E_28: RespInstanceBaseInstanceStatus{
			value: 28,
		}, E_29: RespInstanceBaseInstanceStatus{
			value: 29,
		}, E_30: RespInstanceBaseInstanceStatus{
			value: 30,
		}, E_31: RespInstanceBaseInstanceStatus{
			value: 31,
		}, E_32: RespInstanceBaseInstanceStatus{
			value: 32,
		}, E_33: RespInstanceBaseInstanceStatus{
			value: 33,
		}, E_34: RespInstanceBaseInstanceStatus{
			value: 34,
		}, E_35: RespInstanceBaseInstanceStatus{
			value: 35,
		}, E_36: RespInstanceBaseInstanceStatus{
			value: 36,
		}, E_37: RespInstanceBaseInstanceStatus{
			value: 37,
		}, E_38: RespInstanceBaseInstanceStatus{
			value: 38,
		}, E_39: RespInstanceBaseInstanceStatus{
			value: 39,
		}, E_40: RespInstanceBaseInstanceStatus{
			value: 40,
		}, E_41: RespInstanceBaseInstanceStatus{
			value: 41,
		}, E_42: RespInstanceBaseInstanceStatus{
			value: 42,
		}, E_43: RespInstanceBaseInstanceStatus{
			value: 43,
		}, E_44: RespInstanceBaseInstanceStatus{
			value: 44,
		},
	}
}

func (c RespInstanceBaseInstanceStatus) Value() int32 {
	return c.value
}

func (c RespInstanceBaseInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespInstanceBaseInstanceStatus) UnmarshalJSON(b []byte) error {
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

type RespInstanceBaseSpec struct {
	value string
}

type RespInstanceBaseSpecEnum struct {
	BASIC             RespInstanceBaseSpec
	PROFESSIONAL      RespInstanceBaseSpec
	ENTERPRISE        RespInstanceBaseSpec
	PLATINUM          RespInstanceBaseSpec
	BASIC_IPV6        RespInstanceBaseSpec
	PROFESSIONAL_IPV6 RespInstanceBaseSpec
	ENTERPRISE_IPV6   RespInstanceBaseSpec
	PLATINUM_IPV6     RespInstanceBaseSpec
	PLATINUM_X2       RespInstanceBaseSpec
	PLATINUM_X3       RespInstanceBaseSpec
	PLATINUM_X4       RespInstanceBaseSpec
	PLATINUM_X5       RespInstanceBaseSpec
	PLATINUM_X6       RespInstanceBaseSpec
	PLATINUM_X7       RespInstanceBaseSpec
	PLATINUM_X8       RespInstanceBaseSpec
}

func GetRespInstanceBaseSpecEnum() RespInstanceBaseSpecEnum {
	return RespInstanceBaseSpecEnum{
		BASIC: RespInstanceBaseSpec{
			value: "BASIC",
		},
		PROFESSIONAL: RespInstanceBaseSpec{
			value: "PROFESSIONAL",
		},
		ENTERPRISE: RespInstanceBaseSpec{
			value: "ENTERPRISE",
		},
		PLATINUM: RespInstanceBaseSpec{
			value: "PLATINUM",
		},
		BASIC_IPV6: RespInstanceBaseSpec{
			value: "BASIC_IPV6",
		},
		PROFESSIONAL_IPV6: RespInstanceBaseSpec{
			value: "PROFESSIONAL_IPV6",
		},
		ENTERPRISE_IPV6: RespInstanceBaseSpec{
			value: "ENTERPRISE_IPV6",
		},
		PLATINUM_IPV6: RespInstanceBaseSpec{
			value: "PLATINUM_IPV6",
		},
		PLATINUM_X2: RespInstanceBaseSpec{
			value: "PLATINUM_X2",
		},
		PLATINUM_X3: RespInstanceBaseSpec{
			value: "PLATINUM_X3",
		},
		PLATINUM_X4: RespInstanceBaseSpec{
			value: "PLATINUM_X4",
		},
		PLATINUM_X5: RespInstanceBaseSpec{
			value: "PLATINUM_X5",
		},
		PLATINUM_X6: RespInstanceBaseSpec{
			value: "PLATINUM_X6",
		},
		PLATINUM_X7: RespInstanceBaseSpec{
			value: "PLATINUM_X7",
		},
		PLATINUM_X8: RespInstanceBaseSpec{
			value: "PLATINUM_X8",
		},
	}
}

func (c RespInstanceBaseSpec) Value() string {
	return c.value
}

func (c RespInstanceBaseSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespInstanceBaseSpec) UnmarshalJSON(b []byte) error {
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

type RespInstanceBaseChargingMode struct {
	value int32
}

type RespInstanceBaseChargingModeEnum struct {
	E_0 RespInstanceBaseChargingMode
	E_1 RespInstanceBaseChargingMode
}

func GetRespInstanceBaseChargingModeEnum() RespInstanceBaseChargingModeEnum {
	return RespInstanceBaseChargingModeEnum{
		E_0: RespInstanceBaseChargingMode{
			value: 0,
		}, E_1: RespInstanceBaseChargingMode{
			value: 1,
		},
	}
}

func (c RespInstanceBaseChargingMode) Value() int32 {
	return c.value
}

func (c RespInstanceBaseChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespInstanceBaseChargingMode) UnmarshalJSON(b []byte) error {
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

type RespInstanceBaseLoadbalancerProvider struct {
	value string
}

type RespInstanceBaseLoadbalancerProviderEnum struct {
	LVS RespInstanceBaseLoadbalancerProvider
	ELB RespInstanceBaseLoadbalancerProvider
}

func GetRespInstanceBaseLoadbalancerProviderEnum() RespInstanceBaseLoadbalancerProviderEnum {
	return RespInstanceBaseLoadbalancerProviderEnum{
		LVS: RespInstanceBaseLoadbalancerProvider{
			value: "lvs",
		},
		ELB: RespInstanceBaseLoadbalancerProvider{
			value: "elb",
		},
	}
}

func (c RespInstanceBaseLoadbalancerProvider) Value() string {
	return c.value
}

func (c RespInstanceBaseLoadbalancerProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespInstanceBaseLoadbalancerProvider) UnmarshalJSON(b []byte) error {
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
