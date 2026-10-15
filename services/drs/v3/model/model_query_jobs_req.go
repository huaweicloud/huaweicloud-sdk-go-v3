package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryJobsReq 查询任务列表请求体
type QueryJobsReq struct {

	// **参数解释**：  当前页码，用于分页查询时指定返回第几页的数据。填0可获取全部数据项。  **约束限制**：  - 当per_page填0时，cur_page只能填0或1。  **取值范围**：  不涉及。  **默认取值**：  1
	CurPage *int32 `json:"cur_page,omitempty"`

	// **参数解释**：  每页返回的记录数，用于分页查询时控制每页数据量。  **约束限制**：  - 填0时，cur_page只能填0或1。  **取值范围**：  0-100。  **默认取值**：  10
	PerPage *int32 `json:"per_page,omitempty"`

	// **参数解释**：  迁移任务的业务场景类型。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备  **默认取值**：  不涉及
	DbUseType QueryJobsReqDbUseType `json:"db_use_type"`

	// **参数解释**：  DRS任务引擎类型。  **约束限制**：  不涉及。  **取值范围**：  - mysql：MySQL到MySQL迁移，MySQL到MySQL同步 - mongodb：MongoDB到DDS迁移与同步 - cloudDataGuard-mysql：MySQL到MySQL灾备 - gaussdbv5：GaussDB同步 - mysql-to-kafka：MySQL到Kafka同步 - taurus-to-kafka：TaurusDB到Kafka同步 - gaussdbv5ha-to-kafka：GaussDB集中式版到Kafka同步 - postgresql：PostgreSQL到PostgreSQL同步  **默认取值**：  不涉及
	EngineType *QueryJobsReqEngineType `json:"engine_type,omitempty"`

	// **参数解释**：  企业项目ID，用于按企业项目过滤查询结果。  **约束限制**：  - 不传值时设置为null，不能设置为空字符串。 - 当企业项目开关打开时，可以传值。  **取值范围**：  不涉及。  **默认取值**：  null
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：  任务名称或ID，用于按任务名称或ID过滤查询结果。支持传入多个任务ID，逗号分隔，最多100个。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  任务使用的网络类型。  **约束限制**：  不涉及。  **取值范围**：  - vpn：虚拟专用网络 - vpc：虚拟私有云 - eip：弹性公网IP  **默认取值**：  不涉及
	NetType *QueryJobsReqNetType `json:"net_type,omitempty"`

	// **参数解释**：  服务名称，开启EPS（企业项目服务）时使用。  **约束限制**：  不涉及。  **取值范围**：  - eps：企业项目服务  **默认取值**：  不涉及
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释**：  任务当前运行状态，用于按任务状态过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - STARTJOBING：启动中 - WAITING_FOR_START：等待启动中 - START_JOB_FAILED：启动失败 - PAUSING：已暂停 - FULL_TRANSFER_STARTED：全量开始，灾备场景下为初始化 - FULL_TRANSFER_FAILED：全量失败，灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE：全量完成，灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED：增量开始，灾备场景下为灾备中 - INCRE_TRANSFER_FAILED：增量失败，灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED：结束任务中 - RELEASE_RESOURCE_FAILED：结束任务失败 - RELEASE_RESOURCE_COMPLETE：已结束 - REBUILD_NODE_STARTED：故障恢复中 - REBUILD_NODE_FAILED：故障恢复失败 - CHANGE_JOB_STARTED：任务变更中 - CHANGE_JOB_FAILED：任务变更失败 - DELETED：已删除 - CHILD_TRANSFER_STARTING：再编辑子任务启动中 - CHILD_TRANSFER_STARTED：再编辑子任务迁移中 - CHILD_TRANSFER_COMPLETE：再编辑子任务迁移完成 - CHILD_TRANSFER_FAILED：再编辑子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED：再编辑子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE：再编辑子任务已结束 - NODE_UPGRADE_START：升级开始 - NODE_UPGRADE_COMPLETE：升级完成 - NODE_UPGRADE_FAILED：升级失败  **默认取值**：  不涉及
	Status *QueryJobsReqStatus `json:"status,omitempty"`

	// **参数解释**：  标签信息，用于按标签过滤查询结果，格式为key/value键值对。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Tags map[string]string `json:"tags,omitempty"`

	// **参数解释**：  数据库实例ID列表，用于按数据库实例过滤查询结果。  **约束限制**：  - 最多支持10个实例ID。  **取值范围**：  不涉及。  **默认取值**：  null，表示不使用数据库实例ID过滤
	InstanceIds *[]string `json:"instance_ids,omitempty"`

	// **参数解释**：  DRS绑定的数据库实例IP，用于按数据库实例IP过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  空字符串，表示不使用DRS绑定数据库IP过滤
	InstanceIp *string `json:"instance_ip,omitempty"`
}

func (o QueryJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryJobsReq struct{}"
	}

	return strings.Join([]string{"QueryJobsReq", string(data)}, " ")
}

type QueryJobsReqDbUseType struct {
	value string
}

type QueryJobsReqDbUseTypeEnum struct {
	MIGRATION        QueryJobsReqDbUseType
	SYNC             QueryJobsReqDbUseType
	CLOUD_DATA_GUARD QueryJobsReqDbUseType
}

func GetQueryJobsReqDbUseTypeEnum() QueryJobsReqDbUseTypeEnum {
	return QueryJobsReqDbUseTypeEnum{
		MIGRATION: QueryJobsReqDbUseType{
			value: "migration",
		},
		SYNC: QueryJobsReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: QueryJobsReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c QueryJobsReqDbUseType) Value() string {
	return c.value
}

func (c QueryJobsReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqDbUseType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqEngineType struct {
	value string
}

type QueryJobsReqEngineTypeEnum struct {
	MYSQL                  QueryJobsReqEngineType
	MONGODB                QueryJobsReqEngineType
	CLOUD_DATA_GUARD_MYSQL QueryJobsReqEngineType
	GAUSSDBV5              QueryJobsReqEngineType
	POSTGRESQL             QueryJobsReqEngineType
	MYSQL_TO_KAFKA         QueryJobsReqEngineType
	TAURUS_TO_KAFKA        QueryJobsReqEngineType
	GAUSSDBV5HA_TO_KAFKA   QueryJobsReqEngineType
}

func GetQueryJobsReqEngineTypeEnum() QueryJobsReqEngineTypeEnum {
	return QueryJobsReqEngineTypeEnum{
		MYSQL: QueryJobsReqEngineType{
			value: "mysql",
		},
		MONGODB: QueryJobsReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: QueryJobsReqEngineType{
			value: "cloudDataGuard-mysql",
		},
		GAUSSDBV5: QueryJobsReqEngineType{
			value: "gaussdbv5",
		},
		POSTGRESQL: QueryJobsReqEngineType{
			value: "postgresql",
		},
		MYSQL_TO_KAFKA: QueryJobsReqEngineType{
			value: "mysql-to-kafka",
		},
		TAURUS_TO_KAFKA: QueryJobsReqEngineType{
			value: "taurus-to-kafka",
		},
		GAUSSDBV5HA_TO_KAFKA: QueryJobsReqEngineType{
			value: "gaussdbv5ha-to-kafka",
		},
	}
}

func (c QueryJobsReqEngineType) Value() string {
	return c.value
}

func (c QueryJobsReqEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqEngineType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqNetType struct {
	value string
}

type QueryJobsReqNetTypeEnum struct {
	VPN QueryJobsReqNetType
	VPC QueryJobsReqNetType
	EIP QueryJobsReqNetType
}

func GetQueryJobsReqNetTypeEnum() QueryJobsReqNetTypeEnum {
	return QueryJobsReqNetTypeEnum{
		VPN: QueryJobsReqNetType{
			value: "vpn",
		},
		VPC: QueryJobsReqNetType{
			value: "vpc",
		},
		EIP: QueryJobsReqNetType{
			value: "eip",
		},
	}
}

func (c QueryJobsReqNetType) Value() string {
	return c.value
}

func (c QueryJobsReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqNetType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqStatus struct {
	value string
}

type QueryJobsReqStatusEnum struct {
	CREATING                        QueryJobsReqStatus
	CREATE_FAILED                   QueryJobsReqStatus
	CONFIGURATION                   QueryJobsReqStatus
	STARTJOBING                     QueryJobsReqStatus
	WAITING_FOR_START               QueryJobsReqStatus
	START_JOB_FAILED                QueryJobsReqStatus
	PAUSING                         QueryJobsReqStatus
	FULL_TRANSFER_STARTED           QueryJobsReqStatus
	FULL_TRANSFER_FAILED            QueryJobsReqStatus
	FULL_TRANSFER_COMPLETE          QueryJobsReqStatus
	INCRE_TRANSFER_STARTED          QueryJobsReqStatus
	INCRE_TRANSFER_FAILED           QueryJobsReqStatus
	RELEASE_RESOURCE_STARTED        QueryJobsReqStatus
	RELEASE_RESOURCE_FAILED         QueryJobsReqStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobsReqStatus
	REBUILD_NODE_STARTED            QueryJobsReqStatus
	REBUILD_NODE_FAILED             QueryJobsReqStatus
	CHANGE_JOB_STARTED              QueryJobsReqStatus
	CHANGE_JOB_FAILED               QueryJobsReqStatus
	DELETED                         QueryJobsReqStatus
	CHILD_TRANSFER_STARTING         QueryJobsReqStatus
	CHILD_TRANSFER_STARTED          QueryJobsReqStatus
	CHILD_TRANSFER_COMPLETE         QueryJobsReqStatus
	CHILD_TRANSFER_FAILED           QueryJobsReqStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobsReqStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobsReqStatus
	NODE_UPGRADE_START              QueryJobsReqStatus
	NODE_UPGRADE_COMPLETE           QueryJobsReqStatus
	NODE_UPGRADE_FAILED             QueryJobsReqStatus
}

func GetQueryJobsReqStatusEnum() QueryJobsReqStatusEnum {
	return QueryJobsReqStatusEnum{
		CREATING: QueryJobsReqStatus{
			value: "CREATING",
		},
		CREATE_FAILED: QueryJobsReqStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: QueryJobsReqStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: QueryJobsReqStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: QueryJobsReqStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: QueryJobsReqStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: QueryJobsReqStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: QueryJobsReqStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: QueryJobsReqStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: QueryJobsReqStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: QueryJobsReqStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: QueryJobsReqStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: QueryJobsReqStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: QueryJobsReqStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: QueryJobsReqStatus{
			value: "NODE_UPGRADE_FAILED",
		},
	}
}

func (c QueryJobsReqStatus) Value() string {
	return c.value
}

func (c QueryJobsReqStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqStatus) UnmarshalJSON(b []byte) error {
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
