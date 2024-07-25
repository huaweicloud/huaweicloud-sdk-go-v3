package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DynamicMaskingPolicySet struct {

	// 动态脱敏策略id。
	Id *string `json:"id,omitempty"`

	// 动态脱敏策略名称。英文和汉字开头, 支持英文、汉字、数字、下划线, 2-64字符。
	Name *string `json:"name,omitempty"`

	// 数据源类型 - HIVE数据源 - DWS数据源 - [DLI数据源](tag:nohcs)
	DatasourceType *DynamicMaskingPolicySetDatasourceType `json:"datasource_type,omitempty"`

	// 集群id。请于集群管理页面查看集群ID信息。[当数据源类型为DLI时，该参数需要填写为DLI](tag:nohcs)。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称。请于集群管理页面查看集群名称信息。[当数据源类型为DLI时，该参数需要填写为DLI](tag:nohcs)。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名称。获取方法请参见[获取数据源中的表](getDataTables.html)。
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据表名称, 获取方法请参见[获取数据源中的表](getDataTables.html)。
	TableName *string `json:"table_name,omitempty"`

	// 用户组列表，用户组名称逗号分隔（非必填项，但用户、用户组必须二选其一进行配置）。例如：\"userGroup1,userGroup2\"。
	UserGroups *string `json:"user_groups,omitempty"`

	// 用户列表，用户名称逗号分隔（非必填项，但用户、用户组必须二选其一进行配置），例如：\"user1,user2\"。
	Users *string `json:"users,omitempty"`

	// 同步状态： - UNKNOWN 未知状态 - NOT_SYNC 未同步 - SYNCING 同步中 - SYNC_SUCCESS 同步成功 - SYNC_FAIL 同步失败 - SYNC_PARTIAL_FAIL 存在失败 - DELETE_FAIL 删除失败 - DELETING 删除中 - UPDATING 更新中 - DATA_UPDATED 数据存在更新
	SyncStatus *DynamicMaskingPolicySetSyncStatus `json:"sync_status,omitempty"`

	// 策略同步时间。
	SyncTime *int64 `json:"sync_time,omitempty"`

	// 策略同步日志。
	SyncMsg *string `json:"sync_msg,omitempty"`

	// 策略创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 策略创建者。
	CreateUser *string `json:"create_user,omitempty"`

	// 策略更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 策略更新者。
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o DynamicMaskingPolicySet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DynamicMaskingPolicySet struct{}"
	}

	return strings.Join([]string{"DynamicMaskingPolicySet", string(data)}, " ")
}

type DynamicMaskingPolicySetDatasourceType struct {
	value string
}

type DynamicMaskingPolicySetDatasourceTypeEnum struct {
	HIVE DynamicMaskingPolicySetDatasourceType
	DWS  DynamicMaskingPolicySetDatasourceType
	DLI  DynamicMaskingPolicySetDatasourceType
}

func GetDynamicMaskingPolicySetDatasourceTypeEnum() DynamicMaskingPolicySetDatasourceTypeEnum {
	return DynamicMaskingPolicySetDatasourceTypeEnum{
		HIVE: DynamicMaskingPolicySetDatasourceType{
			value: "HIVE",
		},
		DWS: DynamicMaskingPolicySetDatasourceType{
			value: "DWS",
		},
		DLI: DynamicMaskingPolicySetDatasourceType{
			value: "DLI",
		},
	}
}

func (c DynamicMaskingPolicySetDatasourceType) Value() string {
	return c.value
}

func (c DynamicMaskingPolicySetDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DynamicMaskingPolicySetDatasourceType) UnmarshalJSON(b []byte) error {
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

type DynamicMaskingPolicySetSyncStatus struct {
	value string
}

type DynamicMaskingPolicySetSyncStatusEnum struct {
	UNKNOWN           DynamicMaskingPolicySetSyncStatus
	NOT_SYNC          DynamicMaskingPolicySetSyncStatus
	SYNCING           DynamicMaskingPolicySetSyncStatus
	SYNC_SUCCESS      DynamicMaskingPolicySetSyncStatus
	SYNC_FAIL         DynamicMaskingPolicySetSyncStatus
	SYNC_PARTIAL_FAIL DynamicMaskingPolicySetSyncStatus
	DELETE_FAIL       DynamicMaskingPolicySetSyncStatus
	DELETING          DynamicMaskingPolicySetSyncStatus
	UPDATING          DynamicMaskingPolicySetSyncStatus
	DATA_UPDATED      DynamicMaskingPolicySetSyncStatus
}

func GetDynamicMaskingPolicySetSyncStatusEnum() DynamicMaskingPolicySetSyncStatusEnum {
	return DynamicMaskingPolicySetSyncStatusEnum{
		UNKNOWN: DynamicMaskingPolicySetSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: DynamicMaskingPolicySetSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: DynamicMaskingPolicySetSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: DynamicMaskingPolicySetSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: DynamicMaskingPolicySetSyncStatus{
			value: "SYNC_FAIL",
		},
		SYNC_PARTIAL_FAIL: DynamicMaskingPolicySetSyncStatus{
			value: "SYNC_PARTIAL_FAIL",
		},
		DELETE_FAIL: DynamicMaskingPolicySetSyncStatus{
			value: "DELETE_FAIL",
		},
		DELETING: DynamicMaskingPolicySetSyncStatus{
			value: "DELETING",
		},
		UPDATING: DynamicMaskingPolicySetSyncStatus{
			value: "UPDATING",
		},
		DATA_UPDATED: DynamicMaskingPolicySetSyncStatus{
			value: "DATA_UPDATED",
		},
	}
}

func (c DynamicMaskingPolicySetSyncStatus) Value() string {
	return c.value
}

func (c DynamicMaskingPolicySetSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DynamicMaskingPolicySetSyncStatus) UnmarshalJSON(b []byte) error {
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
