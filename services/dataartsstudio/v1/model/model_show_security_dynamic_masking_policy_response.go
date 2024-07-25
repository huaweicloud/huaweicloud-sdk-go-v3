package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityDynamicMaskingPolicyResponse Response Object
type ShowSecurityDynamicMaskingPolicyResponse struct {

	// 策略id。
	Id *string `json:"id,omitempty"`

	// 策略名称。英文和汉字开头, 支持英文、汉字、数字、下划线, 2-64字符。
	Name *string `json:"name,omitempty"`

	// 数据源类型 - HIVE数据源 - DWS数据源 - [DLI数据源](tag:nohcs)
	DatasourceType *ShowSecurityDynamicMaskingPolicyResponseDatasourceType `json:"datasource_type,omitempty"`

	// 集群id。请于集群管理页面查看集群ID信息。[当数据源类型为DLI时，该参数需要填写为DLI](tag:nohcs)。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称。请于集群管理页面查看集群名称信息。[当数据源类型为DLI时，该参数需要填写为DLI](tag:nohcs)。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名称。获取方法请参见[获取数据源中的表](getDataTables.html)。
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据表id，获取方法请参见[获取数据源中的表](getDataTables.html)。
	TableId *string `json:"table_id,omitempty"`

	// 数据表名称, 获取方法请参见[获取数据源中的表](getDataTables.html)。
	TableName *string `json:"table_name,omitempty"`

	// 用户组列表，用户组名称逗号分隔（非必填项，但用户、用户组必须二选其一进行配置）。例如：\"userGroup1,userGroup2\"。
	UserGroups *string `json:"user_groups,omitempty"`

	// 用户列表，用户名称逗号分隔（非必填项，但用户、用户组必须二选其一进行配置），例如：\"user1,user2\"。
	Users *string `json:"users,omitempty"`

	// 数据连接名称，获取方法请参见[查询数据连接列表](ListDataconnections.html)。
	ConnName *string `json:"conn_name,omitempty"`

	// 数据连接id，获取方法请参见[查询数据连接列表](ListDataconnections.html)。
	ConnId *string `json:"conn_id,omitempty"`

	// 同步状态： - UNKNOWN 未知状态 - NOT_SYNC 未同步 - SYNCING 同步中 - SYNC_SUCCESS 同步成功 - SYNC_FAIL 同步失败 - SYNC_PARTIAL_FAIL 存在失败 - DELETE_FAIL 删除失败 - DELETING 删除中 - UPDATING 更新中 - DATA_UPDATED 数据存在更新
	SyncStatus *ShowSecurityDynamicMaskingPolicyResponseSyncStatus `json:"sync_status,omitempty"`

	// 策略同步信息。
	SyncMsg *string `json:"sync_msg,omitempty"`

	// 同步运行日志, 格式为 字段同步信息+换行符。
	SyncLog *string `json:"sync_log,omitempty"`

	// 策略创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 策略创建者。
	CreateUser *string `json:"create_user,omitempty"`

	// 策略更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 策略更新者。
	UpdateUser *string `json:"update_user,omitempty"`

	// DWS数据源的模式名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 动态数据脱敏策略列表。
	PolicyList     *[]DynamicMaskingPolicy `json:"policy_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowSecurityDynamicMaskingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityDynamicMaskingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityDynamicMaskingPolicyResponse", string(data)}, " ")
}

type ShowSecurityDynamicMaskingPolicyResponseDatasourceType struct {
	value string
}

type ShowSecurityDynamicMaskingPolicyResponseDatasourceTypeEnum struct {
	HIVE ShowSecurityDynamicMaskingPolicyResponseDatasourceType
	DWS  ShowSecurityDynamicMaskingPolicyResponseDatasourceType
	DLI  ShowSecurityDynamicMaskingPolicyResponseDatasourceType
}

func GetShowSecurityDynamicMaskingPolicyResponseDatasourceTypeEnum() ShowSecurityDynamicMaskingPolicyResponseDatasourceTypeEnum {
	return ShowSecurityDynamicMaskingPolicyResponseDatasourceTypeEnum{
		HIVE: ShowSecurityDynamicMaskingPolicyResponseDatasourceType{
			value: "HIVE",
		},
		DWS: ShowSecurityDynamicMaskingPolicyResponseDatasourceType{
			value: "DWS",
		},
		DLI: ShowSecurityDynamicMaskingPolicyResponseDatasourceType{
			value: "DLI",
		},
	}
}

func (c ShowSecurityDynamicMaskingPolicyResponseDatasourceType) Value() string {
	return c.value
}

func (c ShowSecurityDynamicMaskingPolicyResponseDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityDynamicMaskingPolicyResponseDatasourceType) UnmarshalJSON(b []byte) error {
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

type ShowSecurityDynamicMaskingPolicyResponseSyncStatus struct {
	value string
}

type ShowSecurityDynamicMaskingPolicyResponseSyncStatusEnum struct {
	UNKNOWN           ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	NOT_SYNC          ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	SYNCING           ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	SYNC_SUCCESS      ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	SYNC_FAIL         ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	SYNC_PARTIAL_FAIL ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	DELETE_FAIL       ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	DELETING          ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	UPDATING          ShowSecurityDynamicMaskingPolicyResponseSyncStatus
	DATA_UPDATED      ShowSecurityDynamicMaskingPolicyResponseSyncStatus
}

func GetShowSecurityDynamicMaskingPolicyResponseSyncStatusEnum() ShowSecurityDynamicMaskingPolicyResponseSyncStatusEnum {
	return ShowSecurityDynamicMaskingPolicyResponseSyncStatusEnum{
		UNKNOWN: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "SYNC_FAIL",
		},
		SYNC_PARTIAL_FAIL: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "SYNC_PARTIAL_FAIL",
		},
		DELETE_FAIL: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "DELETE_FAIL",
		},
		DELETING: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "DELETING",
		},
		UPDATING: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "UPDATING",
		},
		DATA_UPDATED: ShowSecurityDynamicMaskingPolicyResponseSyncStatus{
			value: "DATA_UPDATED",
		},
	}
}

func (c ShowSecurityDynamicMaskingPolicyResponseSyncStatus) Value() string {
	return c.value
}

func (c ShowSecurityDynamicMaskingPolicyResponseSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityDynamicMaskingPolicyResponseSyncStatus) UnmarshalJSON(b []byte) error {
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
