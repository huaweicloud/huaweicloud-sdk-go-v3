package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportJobsReq 导出任务列表请求体
type ExportJobsReq struct {

	// DRS任务的业务场景类型。用于筛选导出特定业务场景类型的任务列表。 取值： - migration 实时迁移 - sync 实时同步 - backupMigration 备份迁移 - subscription 订阅 - cloudDataGuard 实时灾备
	JobType ExportJobsReqJobType `json:"job_type"`

	// 任务名称。
	Name *string `json:"name,omitempty"`

	// 任务状态。默认查询所有状态的任务 - 当job_type取值：migration,sync,cloudDataGuard     对应的任务状态取值如下：     CREATING：创建中。     CREATE_FAILED：创建失败。     CONFIGURATION：配置中。     STARTJOBING：启动中。     WAITING_FOR_START：等待启动中。     START_JOB_FAILED：任务启动失败。     FULL_TRANSFER_STARTED：全量迁移中，灾备场景为初始化。     FULL_TRANSFER_FAILED：全量迁移失败，灾备场景为初始化失败。     FULL_TRANSFER_COMPLETE：全量迁移完成，灾备场景为初始化完成。     INCRE_TRANSFER_STARTED：增量迁移中，灾备场景为灾备中。     INCRE_TRANSFER_FAILED：增量迁移失败，灾备场景为灾备异常。     RELEASE_RESOURCE_STARTED：结束任务中。     RELEASE_RESOURCE_FAILED：结束任务失败。     RELEASE_RESOURCE_COMPLETE：已结束。     CHANGE_JOB_STARTED：任务变更中。     CHANGE_JOB_FAILED：任务变更失败。     CHILD_TRANSFER_STARTING：子任务启动中。     CHILD_TRANSFER_STARTED：子任务迁移中。     CHILD_TRANSFER_COMPLETE：子任务迁移完成。     CHILD_TRANSFER_FAILED：子任务迁移失败。     RELEASE_CHILD_TRANSFER_STARTED：子任务结束中。     RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束。其中，异常状态可单独查询，也可以通过以下方式查询全部异常任务：CREATE_FAILED，START_JOB_FAILED，FULL_TRANSFER_FAILED，INCRE_TRANSFER_FAILED，RELEASE_RESOURCE_FAILED，CHANGE_JOB_FAILED，CHILD_TRANSFER_FAILED。 - 当job_type取值：backupMigration   对应的任务状态取值如下：     SUBSCRIPTION_STARTED：创建中。     CREATE_FAILED：创建失败。     SUBSCRIPTION_FAILED：订阅失败。     CONFIGURATION：配置中。     STARTJOBING：启动中。     REBUILD_NODE_STARTED：故障恢复。     REBUILD_NODE_FAILED：故障恢复失败。     START_JOB_FAILED：任务启动失败。 - 当job_type取值：subscription   对应的任务状态取值如下：     SUCCESS：成功。     TRANSFERRING：备份迁移中。     FAILED：失败。     PRECHECK FAILED：预检查失败。
	Status *ExportJobsReqStatus `json:"status,omitempty"`

	// 数据库实例ID，最大数量为10。
	DbsInstanceIds *[]string `json:"dbs_instance_ids,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 创建时间。 当job_type取值：migration,sync,cloudDataGuard， 任务创建时间，比如\"2026-05-25T13:41:00.127Z|2026-05-28T13:41:00.999Z\" 当job_type取值：backupMigration， 表示任务创建时间，比如\"2026-05-25T13:41:00.127Z|2026-05-28T13:41:00.999Z\" 当job_type取值：subscription，表示消费时间点，比如\"2026-05-25T13:25:46.383Z|2026-05-28T13:25:46.999Z\"
	CreateAt *string `json:"create_at,omitempty"`

	// 完成时间。在job_type是backupMigration时支持该参数填写， 表示任务完成时间，比如\"2026-05-25T13:41:00.127Z|2026-05-28T13:41:00.999Z\"
	CompletedAt *string `json:"completed_at,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 引擎类型。当job_type取值：migration,sync,cloudDataGuard，该参数支持填写。当job_type取值：backupMigration 时，该参数默认值是sqlserver
	EngineType *string `json:"engine_type,omitempty"`

	// 网络类型。取值： eip：公网网络。 vpc：VPC网络。 vpn：VPN、专线网络。 当job_type取值：migration,sync,cloudDataGuard，该参数支持填写。
	NetType *ExportJobsReqNetType `json:"net_type,omitempty"`

	// 计费标识。当job_type取值：migration,sync,cloudDataGuard,subscription，该参数支持填写。取值\"0\"表示计费中，\"1\"表示未计费
	BillingTag *ExportJobsReqBillingTag `json:"billing_tag,omitempty"`

	// 计费模式。当job_type取值：migration,sync,cloudDataGuard，支持该参数填写。取值：\"1\"表示包年包月，\"0\"表示按需计费
	BillingMode *ExportJobsReqBillingMode `json:"billing_mode,omitempty"`

	// 实例绑定的公网IP。当job_type取值：migration,sync,cloudDataGuard，支持该参数填写。
	PublicIp *string `json:"public_ip,omitempty"`

	// DRS绑定的数据库实例IP。当job_type取值：migration,sync,cloudDataGuard，支持该参数填写。
	InstanceIp *string `json:"instance_ip,omitempty"`

	// DRS绑定的内网IP。当job_type取值：migration,sync,cloudDataGuard，支持该参数填写。
	InnerIp *string `json:"inner_ip,omitempty"`

	// 规格类型。当job_type取值：migration,sync,cloudDataGuard，支持该参数填写。 取值如下： 2xlarge：极大规格。 xlarge：超大规格。 high：大规格。 medium：中规格。 small：小规格。 micro：极小规格。
	SpecType *ExportJobsReqSpecType `json:"spec_type,omitempty"`

	// 数据流动方向。 当job_type取值migration,sync时支持传值。 取值： - up：入云 - down：出云 - non-dbs：自建 - multiWrite：双向同步。 当job_type取值cloudDataGuard时支持传值， 取值： - up：本云为备 - down：本云为主 - multiWrite：双主灾备
	Direction *ExportJobsReqDirection `json:"direction,omitempty"`

	// 同步类型。当job_type取值migration,sync时支持传值。
	TaskType *ExportJobsReqTaskType `json:"task_type,omitempty"`

	// 返回结果按该关键字排序，默认为“create_time”。 当前支持排序的关键字： name status create_time net_type job_direction pay_mode
	SortKey *string `json:"sort_key,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为“desc”）。
	SortDir *string `json:"sort_dir,omitempty"`

	// 标签
	Tag *string `json:"tag,omitempty"`
}

func (o ExportJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobsReq struct{}"
	}

	return strings.Join([]string{"ExportJobsReq", string(data)}, " ")
}

type ExportJobsReqJobType struct {
	value string
}

type ExportJobsReqJobTypeEnum struct {
	MIGRATION        ExportJobsReqJobType
	SYNC             ExportJobsReqJobType
	CLOUD_DATA_GUARD ExportJobsReqJobType
	BACKUP_MIGRATION ExportJobsReqJobType
	SUBSCRIPTION     ExportJobsReqJobType
}

func GetExportJobsReqJobTypeEnum() ExportJobsReqJobTypeEnum {
	return ExportJobsReqJobTypeEnum{
		MIGRATION: ExportJobsReqJobType{
			value: "migration",
		},
		SYNC: ExportJobsReqJobType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ExportJobsReqJobType{
			value: "cloudDataGuard",
		},
		BACKUP_MIGRATION: ExportJobsReqJobType{
			value: "backupMigration",
		},
		SUBSCRIPTION: ExportJobsReqJobType{
			value: "subscription",
		},
	}
}

func (c ExportJobsReqJobType) Value() string {
	return c.value
}

func (c ExportJobsReqJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqJobType) UnmarshalJSON(b []byte) error {
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

type ExportJobsReqStatus struct {
	value string
}

type ExportJobsReqStatusEnum struct {
	CREATING                        ExportJobsReqStatus
	CREATE_FAILED                   ExportJobsReqStatus
	CONFIGURATION                   ExportJobsReqStatus
	STARTJOBING                     ExportJobsReqStatus
	WAITING_FOR_START               ExportJobsReqStatus
	START_JOB_FAILED                ExportJobsReqStatus
	FULL_TRANSFER_STARTED           ExportJobsReqStatus
	FULL_TRANSFER_FAILED            ExportJobsReqStatus
	FULL_TRANSFER_COMPLETE          ExportJobsReqStatus
	INCRE_TRANSFER_STARTED          ExportJobsReqStatus
	INCRE_TRANSFER_FAILED           ExportJobsReqStatus
	RELEASE_RESOURCE_STARTED        ExportJobsReqStatus
	RELEASE_RESOURCE_FAILED         ExportJobsReqStatus
	RELEASE_RESOURCE_COMPLETE       ExportJobsReqStatus
	CHANGE_JOB_STARTED              ExportJobsReqStatus
	CHANGE_JOB_FAILED               ExportJobsReqStatus
	CHILD_TRANSFER_STARTING         ExportJobsReqStatus
	CHILD_TRANSFER_STARTED          ExportJobsReqStatus
	CHILD_TRANSFER_COMPLETE         ExportJobsReqStatus
	CHILD_TRANSFER_FAILED           ExportJobsReqStatus
	RELEASE_CHILD_TRANSFER_STARTED  ExportJobsReqStatus
	RELEASE_CHILD_TRANSFER_COMPLETE ExportJobsReqStatus
	SUBSCRIPTION_STARTED            ExportJobsReqStatus
	SUBSCRIPTION_FAILED             ExportJobsReqStatus
	EMPTY                           ExportJobsReqStatus
	REBUILD_NODE_FAILED             ExportJobsReqStatus
	SUCCESS                         ExportJobsReqStatus
	TRANSFERRING                    ExportJobsReqStatus
	FAILED                          ExportJobsReqStatus
	PRECHECK_FAILED                 ExportJobsReqStatus
}

func GetExportJobsReqStatusEnum() ExportJobsReqStatusEnum {
	return ExportJobsReqStatusEnum{
		CREATING: ExportJobsReqStatus{
			value: "CREATING",
		},
		CREATE_FAILED: ExportJobsReqStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: ExportJobsReqStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: ExportJobsReqStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: ExportJobsReqStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: ExportJobsReqStatus{
			value: "START_JOB_FAILED",
		},
		FULL_TRANSFER_STARTED: ExportJobsReqStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: ExportJobsReqStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: ExportJobsReqStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: ExportJobsReqStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: ExportJobsReqStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: ExportJobsReqStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: ExportJobsReqStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: ExportJobsReqStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		CHANGE_JOB_STARTED: ExportJobsReqStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: ExportJobsReqStatus{
			value: "CHANGE_JOB_FAILED",
		},
		CHILD_TRANSFER_STARTING: ExportJobsReqStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: ExportJobsReqStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: ExportJobsReqStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: ExportJobsReqStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: ExportJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: ExportJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		SUBSCRIPTION_STARTED: ExportJobsReqStatus{
			value: "SUBSCRIPTION_STARTED",
		},
		SUBSCRIPTION_FAILED: ExportJobsReqStatus{
			value: "SUBSCRIPTION_FAILED",
		},
		EMPTY: ExportJobsReqStatus{
			value: "",
		},
		REBUILD_NODE_FAILED: ExportJobsReqStatus{
			value: "REBUILD_NODE_FAILED",
		},
		SUCCESS: ExportJobsReqStatus{
			value: "SUCCESS",
		},
		TRANSFERRING: ExportJobsReqStatus{
			value: "TRANSFERRING",
		},
		FAILED: ExportJobsReqStatus{
			value: "FAILED",
		},
		PRECHECK_FAILED: ExportJobsReqStatus{
			value: "PRECHECK FAILED",
		},
	}
}

func (c ExportJobsReqStatus) Value() string {
	return c.value
}

func (c ExportJobsReqStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqStatus) UnmarshalJSON(b []byte) error {
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

type ExportJobsReqNetType struct {
	value string
}

type ExportJobsReqNetTypeEnum struct {
	EIP ExportJobsReqNetType
	VPC ExportJobsReqNetType
	VPN ExportJobsReqNetType
}

func GetExportJobsReqNetTypeEnum() ExportJobsReqNetTypeEnum {
	return ExportJobsReqNetTypeEnum{
		EIP: ExportJobsReqNetType{
			value: "eip",
		},
		VPC: ExportJobsReqNetType{
			value: "vpc",
		},
		VPN: ExportJobsReqNetType{
			value: "vpn",
		},
	}
}

func (c ExportJobsReqNetType) Value() string {
	return c.value
}

func (c ExportJobsReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqNetType) UnmarshalJSON(b []byte) error {
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

type ExportJobsReqBillingTag struct {
	value string
}

type ExportJobsReqBillingTagEnum struct {
	E_0 ExportJobsReqBillingTag
	E_1 ExportJobsReqBillingTag
}

func GetExportJobsReqBillingTagEnum() ExportJobsReqBillingTagEnum {
	return ExportJobsReqBillingTagEnum{
		E_0: ExportJobsReqBillingTag{
			value: "0",
		},
		E_1: ExportJobsReqBillingTag{
			value: "1",
		},
	}
}

func (c ExportJobsReqBillingTag) Value() string {
	return c.value
}

func (c ExportJobsReqBillingTag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqBillingTag) UnmarshalJSON(b []byte) error {
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

type ExportJobsReqBillingMode struct {
	value string
}

type ExportJobsReqBillingModeEnum struct {
	E_0 ExportJobsReqBillingMode
	E_1 ExportJobsReqBillingMode
}

func GetExportJobsReqBillingModeEnum() ExportJobsReqBillingModeEnum {
	return ExportJobsReqBillingModeEnum{
		E_0: ExportJobsReqBillingMode{
			value: "0",
		},
		E_1: ExportJobsReqBillingMode{
			value: "1",
		},
	}
}

func (c ExportJobsReqBillingMode) Value() string {
	return c.value
}

func (c ExportJobsReqBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqBillingMode) UnmarshalJSON(b []byte) error {
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

type ExportJobsReqSpecType struct {
	value string
}

type ExportJobsReqSpecTypeEnum struct {
	E_2XLARGE ExportJobsReqSpecType
	XLARGE    ExportJobsReqSpecType
	HIGH      ExportJobsReqSpecType
	MEDIUM    ExportJobsReqSpecType
	SMALL     ExportJobsReqSpecType
	MICRO     ExportJobsReqSpecType
}

func GetExportJobsReqSpecTypeEnum() ExportJobsReqSpecTypeEnum {
	return ExportJobsReqSpecTypeEnum{
		E_2XLARGE: ExportJobsReqSpecType{
			value: "2xlarge",
		},
		XLARGE: ExportJobsReqSpecType{
			value: "xlarge",
		},
		HIGH: ExportJobsReqSpecType{
			value: "high",
		},
		MEDIUM: ExportJobsReqSpecType{
			value: "medium",
		},
		SMALL: ExportJobsReqSpecType{
			value: "small",
		},
		MICRO: ExportJobsReqSpecType{
			value: "micro",
		},
	}
}

func (c ExportJobsReqSpecType) Value() string {
	return c.value
}

func (c ExportJobsReqSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqSpecType) UnmarshalJSON(b []byte) error {
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

type ExportJobsReqDirection struct {
	value string
}

type ExportJobsReqDirectionEnum struct {
	UP          ExportJobsReqDirection
	DOWN        ExportJobsReqDirection
	NON_DBS     ExportJobsReqDirection
	MULTI_WRITE ExportJobsReqDirection
}

func GetExportJobsReqDirectionEnum() ExportJobsReqDirectionEnum {
	return ExportJobsReqDirectionEnum{
		UP: ExportJobsReqDirection{
			value: "up",
		},
		DOWN: ExportJobsReqDirection{
			value: "down",
		},
		NON_DBS: ExportJobsReqDirection{
			value: "non-dbs",
		},
		MULTI_WRITE: ExportJobsReqDirection{
			value: "multiWrite",
		},
	}
}

func (c ExportJobsReqDirection) Value() string {
	return c.value
}

func (c ExportJobsReqDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqDirection) UnmarshalJSON(b []byte) error {
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

type ExportJobsReqTaskType struct {
	value string
}

type ExportJobsReqTaskTypeEnum struct {
	FULL_INCR_TRANS ExportJobsReqTaskType
	FULL_TRANS      ExportJobsReqTaskType
	INCR_TRANS      ExportJobsReqTaskType
}

func GetExportJobsReqTaskTypeEnum() ExportJobsReqTaskTypeEnum {
	return ExportJobsReqTaskTypeEnum{
		FULL_INCR_TRANS: ExportJobsReqTaskType{
			value: "FULL_INCR_TRANS",
		},
		FULL_TRANS: ExportJobsReqTaskType{
			value: "FULL_TRANS",
		},
		INCR_TRANS: ExportJobsReqTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c ExportJobsReqTaskType) Value() string {
	return c.value
}

func (c ExportJobsReqTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobsReqTaskType) UnmarshalJSON(b []byte) error {
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
