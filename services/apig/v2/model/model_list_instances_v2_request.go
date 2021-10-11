package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListInstancesV2Request struct {
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
	// 实例编号

	InstanceId *string `json:"instance_id,omitempty"`
	// 实例名称

	InstanceName *string `json:"instance_name,omitempty"`
	// 实例状态： - Creating：创建中 - CreateSuccess：创建成功 - CreateFail：创建失败 - Initing：初始化中 - Registering：注册中 - Running：运行中 - InitingFailed：初始化失败 - RegisterFailed：注册失败 - Installing：安装中 - InstallFailed：安装失败 - Updating：升级中 - UpdateFailed：升级失败 - Rollbacking：回滚中 - RollbackSuccess：回滚成功 - RollbackFailed：回滚失败 - Deleting：删除中 - DeleteFailed：删除失败 - Unregistering：注销中 - UnRegisterFailed：注销失败 - CreateTimeout：创建超时 - InitTimeout：初始化超时 - RegisterTimeout：注册超时 - InstallTimeout：安装超时 - UpdateTimeout：升级超时 - RollbackTimeout：回滚超时 - DeleteTimeout：删除超时 - UnregisterTimeout：注销超时 - Starting：启动中 - Freezing：冻结中 - Frozen：已冻结 - Restarting：重启中 - RestartFail：重启失败 - Unhealthy：实例异常 - RestartTimeout：重启超时

	Status *ListInstancesV2RequestStatus `json:"status,omitempty"`
}

func (o ListInstancesV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesV2Request struct{}"
	}

	return strings.Join([]string{"ListInstancesV2Request", string(data)}, " ")
}

type ListInstancesV2RequestStatus struct {
	value string
}

type ListInstancesV2RequestStatusEnum struct {
	CREATING           ListInstancesV2RequestStatus
	CREATE_SUCCESS     ListInstancesV2RequestStatus
	CREATE_FAIL        ListInstancesV2RequestStatus
	INITING            ListInstancesV2RequestStatus
	REGISTERING        ListInstancesV2RequestStatus
	RUNNING            ListInstancesV2RequestStatus
	INITING_FAILED     ListInstancesV2RequestStatus
	REGISTER_FAILED    ListInstancesV2RequestStatus
	INSTALLING         ListInstancesV2RequestStatus
	INSTALL_FAILED     ListInstancesV2RequestStatus
	UPDATING           ListInstancesV2RequestStatus
	UPDATE_FAILED      ListInstancesV2RequestStatus
	ROLLBACKING        ListInstancesV2RequestStatus
	ROLLBACK_SUCCESS   ListInstancesV2RequestStatus
	ROLLBACK_FAILED    ListInstancesV2RequestStatus
	DELETING           ListInstancesV2RequestStatus
	DELETE_FAILED      ListInstancesV2RequestStatus
	UNREGISTERING      ListInstancesV2RequestStatus
	UN_REGISTER_FAILED ListInstancesV2RequestStatus
	CREATE_TIMEOUT     ListInstancesV2RequestStatus
	INIT_TIMEOUT       ListInstancesV2RequestStatus
	REGISTER_TIMEOUT   ListInstancesV2RequestStatus
	INSTALL_TIMEOUT    ListInstancesV2RequestStatus
	UPDATE_TIMEOUT     ListInstancesV2RequestStatus
	ROLLBACK_TIMEOUT   ListInstancesV2RequestStatus
	DELETE_TIMEOUT     ListInstancesV2RequestStatus
	UNREGISTER_TIMEOUT ListInstancesV2RequestStatus
	STARTING           ListInstancesV2RequestStatus
	FREEZING           ListInstancesV2RequestStatus
	FROZEN             ListInstancesV2RequestStatus
	RESTARTING         ListInstancesV2RequestStatus
	RESTART_FAIL       ListInstancesV2RequestStatus
	UNHEALTHY          ListInstancesV2RequestStatus
	RESTART_TIMEOUT    ListInstancesV2RequestStatus
}

func GetListInstancesV2RequestStatusEnum() ListInstancesV2RequestStatusEnum {
	return ListInstancesV2RequestStatusEnum{
		CREATING: ListInstancesV2RequestStatus{
			value: "Creating",
		},
		CREATE_SUCCESS: ListInstancesV2RequestStatus{
			value: "CreateSuccess",
		},
		CREATE_FAIL: ListInstancesV2RequestStatus{
			value: "CreateFail",
		},
		INITING: ListInstancesV2RequestStatus{
			value: "Initing",
		},
		REGISTERING: ListInstancesV2RequestStatus{
			value: "Registering",
		},
		RUNNING: ListInstancesV2RequestStatus{
			value: "Running",
		},
		INITING_FAILED: ListInstancesV2RequestStatus{
			value: "InitingFailed",
		},
		REGISTER_FAILED: ListInstancesV2RequestStatus{
			value: "RegisterFailed",
		},
		INSTALLING: ListInstancesV2RequestStatus{
			value: "Installing",
		},
		INSTALL_FAILED: ListInstancesV2RequestStatus{
			value: "InstallFailed",
		},
		UPDATING: ListInstancesV2RequestStatus{
			value: "Updating",
		},
		UPDATE_FAILED: ListInstancesV2RequestStatus{
			value: "UpdateFailed",
		},
		ROLLBACKING: ListInstancesV2RequestStatus{
			value: "Rollbacking",
		},
		ROLLBACK_SUCCESS: ListInstancesV2RequestStatus{
			value: "RollbackSuccess",
		},
		ROLLBACK_FAILED: ListInstancesV2RequestStatus{
			value: "RollbackFailed",
		},
		DELETING: ListInstancesV2RequestStatus{
			value: "Deleting",
		},
		DELETE_FAILED: ListInstancesV2RequestStatus{
			value: "DeleteFailed",
		},
		UNREGISTERING: ListInstancesV2RequestStatus{
			value: "Unregistering",
		},
		UN_REGISTER_FAILED: ListInstancesV2RequestStatus{
			value: "UnRegisterFailed",
		},
		CREATE_TIMEOUT: ListInstancesV2RequestStatus{
			value: "CreateTimeout",
		},
		INIT_TIMEOUT: ListInstancesV2RequestStatus{
			value: "InitTimeout",
		},
		REGISTER_TIMEOUT: ListInstancesV2RequestStatus{
			value: "RegisterTimeout",
		},
		INSTALL_TIMEOUT: ListInstancesV2RequestStatus{
			value: "InstallTimeout",
		},
		UPDATE_TIMEOUT: ListInstancesV2RequestStatus{
			value: "UpdateTimeout",
		},
		ROLLBACK_TIMEOUT: ListInstancesV2RequestStatus{
			value: "RollbackTimeout",
		},
		DELETE_TIMEOUT: ListInstancesV2RequestStatus{
			value: "DeleteTimeout",
		},
		UNREGISTER_TIMEOUT: ListInstancesV2RequestStatus{
			value: "UnregisterTimeout",
		},
		STARTING: ListInstancesV2RequestStatus{
			value: "Starting",
		},
		FREEZING: ListInstancesV2RequestStatus{
			value: "Freezing",
		},
		FROZEN: ListInstancesV2RequestStatus{
			value: "Frozen",
		},
		RESTARTING: ListInstancesV2RequestStatus{
			value: "Restarting",
		},
		RESTART_FAIL: ListInstancesV2RequestStatus{
			value: "RestartFail",
		},
		UNHEALTHY: ListInstancesV2RequestStatus{
			value: "Unhealthy",
		},
		RESTART_TIMEOUT: ListInstancesV2RequestStatus{
			value: "RestartTimeout",
		},
	}
}

func (c ListInstancesV2RequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListInstancesV2RequestStatus) UnmarshalJSON(b []byte) error {
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
