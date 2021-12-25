package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RespInstanceBase struct {
	// 实例编号

	Id *string `json:"id,omitempty"`
	// 实例所属项目编号

	ProjectId *string `json:"project_id,omitempty"`
	// 实例名称

	InstanceName *string `json:"instance_name,omitempty"`
	// 实例状态： - Creating：创建中 - CreateSuccess：创建成功 - CreateFail：创建失败 - Initing：初始化中 - Registering：注册中 - Running：运行中 - InitingFailed：初始化失败 - RegisterFailed：注册失败 - Installing：安装中 - InstallFailed：安装失败 - Updating：升级中 - UpdateFailed：升级失败 - Rollbacking：回滚中 - RollbackSuccess：回滚成功 - RollbackFailed：回滚失败 - Deleting：删除中 - DeleteFailed：删除失败 - Unregistering：注销中 - UnRegisterFailed：注销失败 - CreateTimeout：创建超时 - InitTimeout：初始化超时 - RegisterTimeout：注册超时 - InstallTimeout：安装超时 - UpdateTimeout：升级超时 - RollbackTimeout：回滚超时 - DeleteTimeout：删除超时 - UnregisterTimeout：注销超时 - Starting：启动中 - Freezing：冻结中 - Frozen：已冻结 - Restarting：重启中 - RestartFail：重启失败 - Unhealthy：实例异常 - RestartTimeout：重启超时

	Status *RespInstanceBaseStatus `json:"status,omitempty"`
	// 实例状态对应编号 - 1：创建中 - 2：创建成功 - 3：创建失败 - 4：初始化中 - 5：注册中 - 6：运行中 - 7：初始化失败 - 8：注册失败 - 10：安装中 - 11：安装失败 - 12：升级中 - 13：升级失败 - 20：回滚中 - 21：回滚成功 - 22：回滚失败 - 23：删除中 - 24：删除失败 - 25：注销中 - 26：注销失败 - 27：创建超时 - 28：初始化超时 - 29：注册超时 - 30：安装超时 - 31：升级超时 - 32：回滚超时 - 33：删除超时 - 34：注销超时 - 35：启动中 - 36：冻结中 - 37：已冻结 - 38：重启中 - 39：重启失败 - 40：实例异常 - 41：重启超时

	InstanceStatus *int32 `json:"instance_status,omitempty"`
	// 实例类型  暂不支持

	Type *string `json:"type,omitempty"`
	// 实例规格： - ROMA_BASIC：基础版实例 - ROMA_PROFESSIONAL：专业版实例 - ROMA_ENTERPRISE：企业版实例 - ROMA_PLATINUM：铂金版实例 [- ROMA_BASIC_IPV6：基础版IPV6实例](tag:hcs) [- ROMA_PROFESSIONAL_IPV6：专业版IPV6实例](tag:hcs) [- ROMA_ENTERPRISE_IPV6：企业版IPV6实例](tag:hcs) [- ROMA_PLATINUM_IPV6：铂金版IPV6实例](tag:hcs)

	Spec *RespInstanceBaseSpec `json:"spec,omitempty"`
	// 实例创建时间。unix时间戳格式。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 企业项目ID，企业帐号必填

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 实例绑定的弹性IP地址

	EipAddress *string `json:"eip_address,omitempty"`
	// 实例计费方式： - 0：按需计费 - 1：包周期计费

	ChargingMode *int32 `json:"charging_mode,omitempty"`
	// 包周期计费订单编号

	CbcMetadata *string `json:"cbc_metadata,omitempty"`
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
	}
}

func (c RespInstanceBaseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespInstanceBaseStatus) UnmarshalJSON(b []byte) error {
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

type RespInstanceBaseSpec struct {
	value string
}

type RespInstanceBaseSpecEnum struct {
	ROMA_BASIC             RespInstanceBaseSpec
	ROMA_PROFESSIONAL      RespInstanceBaseSpec
	ROMA_ENTERPRISE        RespInstanceBaseSpec
	ROMA_PLATINUM          RespInstanceBaseSpec
	ROMA_BASIC_IPV6        RespInstanceBaseSpec
	ROMA_PROFESSIONAL_IPV6 RespInstanceBaseSpec
	ROMA_ENTERPRISE_IPV6   RespInstanceBaseSpec
	ROMA_PLATINUM_IPV6     RespInstanceBaseSpec
}

func GetRespInstanceBaseSpecEnum() RespInstanceBaseSpecEnum {
	return RespInstanceBaseSpecEnum{
		ROMA_BASIC: RespInstanceBaseSpec{
			value: "ROMA_BASIC",
		},
		ROMA_PROFESSIONAL: RespInstanceBaseSpec{
			value: "ROMA_PROFESSIONAL",
		},
		ROMA_ENTERPRISE: RespInstanceBaseSpec{
			value: "ROMA_ENTERPRISE",
		},
		ROMA_PLATINUM: RespInstanceBaseSpec{
			value: "ROMA_PLATINUM",
		},
		ROMA_BASIC_IPV6: RespInstanceBaseSpec{
			value: "ROMA_BASIC_IPV6",
		},
		ROMA_PROFESSIONAL_IPV6: RespInstanceBaseSpec{
			value: "ROMA_PROFESSIONAL_IPV6",
		},
		ROMA_ENTERPRISE_IPV6: RespInstanceBaseSpec{
			value: "ROMA_ENTERPRISE_IPV6",
		},
		ROMA_PLATINUM_IPV6: RespInstanceBaseSpec{
			value: "ROMA_PLATINUM_IPV6",
		},
	}
}

func (c RespInstanceBaseSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RespInstanceBaseSpec) UnmarshalJSON(b []byte) error {
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
