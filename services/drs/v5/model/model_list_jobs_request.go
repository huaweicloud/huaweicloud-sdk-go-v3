package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us：英文 - zh-cn：中文  **默认取值**：  en-us
	XLanguage *ListJobsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  DRS任务场景类型，用于区分不同的数据复制场景。  **约束限制**：  不涉及。  **取值范围**：  - migration：实时迁移 - sync：实时同步 - cloudDataGuard：实时灾备 - replay：流量回放  **默认取值**：  不涉及
	JobType ListJobsRequestJobType `json:"job_type"`

	// **参数解释**：  任务ID或名称，用于按任务ID或名称过滤查询结果。  **约束限制**：  - 支持输入多个任务ID，以逗号分隔，最多10个。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  任务当前运行状态，用于按任务状态过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - STARTJOBING：启动中 - WAITING_FOR_START：等待启动中 - START_JOB_FAILED：启动失败 - FULL_TRANSFER_STARTED：全量开始，灾备场景下为初始化 - FULL_TRANSFER_FAILED：全量失败，灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE：全量完成，灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED：增量开始，灾备场景下为灾备中 - INCRE_TRANSFER_FAILED：增量失败，灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED：结束任务中 - RELEASE_RESOURCE_FAILED：结束任务失败 - RELEASE_RESOURCE_COMPLETE：已结束 - CHANGE_JOB_STARTED：任务变更中 - CHANGE_JOB_FAILED：任务变更失败 - CHILD_TRANSFER_STARTING：子任务启动中 - CHILD_TRANSFER_STARTED：子任务迁移中 - CHILD_TRANSFER_COMPLETE：子任务迁移完成 - CHILD_TRANSFER_FAILED：子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED：子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束  **默认取值**：  不涉及
	Status *ListJobsRequestStatus `json:"status,omitempty"`

	// **参数解释**：  DRS任务引擎类型，用于按引擎类型过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用 - redis-to-gaussredis：Redis迁移到GeminiDB Redis，实时迁移场景使用 - rediscluster-to-gaussredis：Redis集群迁移到GeminiDB Redis，实时迁移场景使用  **默认取值**：  不涉及
	EngineType *ListJobsRequestEngineType `json:"engine_type,omitempty"`

	// **参数解释**：  任务使用的网络类型，用于按网络类型过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  - eip：公网网络 - vpc：VPC网络，灾备场景不支持选择VPC网络 - vpn：VPN、专线网络  **默认取值**：  不涉及
	NetType *ListJobsRequestNetType `json:"net_type,omitempty"`

	// **参数解释**：  企业项目ID，用于按企业项目过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  空字符串，表示查询所有企业项目任务
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：  偏移量，表示查询该偏移量后面的记录，用于分页查询。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询返回记录的数量限制，用于分页查询时控制每页返回的数据条数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  返回结果的排序关键字，指定任务列表的排序依据。  **约束限制**：  不涉及。  **取值范围**：  - name：按任务名称排序 - status：按任务状态排序 - create_time：按创建时间排序 - net_type：按网络类型排序 - job_direction：按任务方向排序 - pay_mode：按付费模式排序  **默认取值**：  create_time
	SortKey *ListJobsRequestSortKey `json:"sort_key,omitempty"`

	// **参数解释**：  排序方向，指定返回结果按升序或降序排列。  **约束限制**：  不涉及。  **取值范围**：  - desc：降序 - asc：升序  **默认取值**：  desc
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**：  数据库实例ID列表，用于按数据库实例过滤查询结果。  **约束限制**：  - 最多支持10个实例ID。  **取值范围**：  不涉及。  **默认取值**：  null，表示不使用数据库实例ID过滤
	InstanceIds *[]string `json:"instance_ids,omitempty"`

	// **参数解释**：  DRS绑定的数据库实例IP，用于按数据库实例IP过滤查询结果。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  空字符串，表示不使用DRS绑定数据库IP过滤
	InstanceIp *string `json:"instance_ip,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}

type ListJobsRequestXLanguage struct {
	value string
}

type ListJobsRequestXLanguageEnum struct {
	EN_US ListJobsRequestXLanguage
	ZH_CN ListJobsRequestXLanguage
}

func GetListJobsRequestXLanguageEnum() ListJobsRequestXLanguageEnum {
	return ListJobsRequestXLanguageEnum{
		EN_US: ListJobsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListJobsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListJobsRequestXLanguage) Value() string {
	return c.value
}

func (c ListJobsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestJobType struct {
	value string
}

type ListJobsRequestJobTypeEnum struct {
	MIGRATION        ListJobsRequestJobType
	SYNC             ListJobsRequestJobType
	CLOUD_DATA_GUARD ListJobsRequestJobType
	REPLAY           ListJobsRequestJobType
}

func GetListJobsRequestJobTypeEnum() ListJobsRequestJobTypeEnum {
	return ListJobsRequestJobTypeEnum{
		MIGRATION: ListJobsRequestJobType{
			value: "migration",
		},
		SYNC: ListJobsRequestJobType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ListJobsRequestJobType{
			value: "cloudDataGuard",
		},
		REPLAY: ListJobsRequestJobType{
			value: "replay",
		},
	}
}

func (c ListJobsRequestJobType) Value() string {
	return c.value
}

func (c ListJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestJobType) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestStatus struct {
	value string
}

type ListJobsRequestStatusEnum struct {
	CREATING                        ListJobsRequestStatus
	CREATE_FAILED                   ListJobsRequestStatus
	CONFIGURATION                   ListJobsRequestStatus
	STARTJOBING                     ListJobsRequestStatus
	WAITING_FOR_START               ListJobsRequestStatus
	START_JOB_FAILED                ListJobsRequestStatus
	FULL_TRANSFER_STARTED           ListJobsRequestStatus
	FULL_TRANSFER_FAILED            ListJobsRequestStatus
	FULL_TRANSFER_COMPLETE          ListJobsRequestStatus
	INCRE_TRANSFER_STARTED          ListJobsRequestStatus
	INCRE_TRANSFER_FAILED           ListJobsRequestStatus
	RELEASE_RESOURCE_STARTED        ListJobsRequestStatus
	RELEASE_RESOURCE_FAILED         ListJobsRequestStatus
	RELEASE_RESOURCE_COMPLETE       ListJobsRequestStatus
	CHANGE_JOB_STARTED              ListJobsRequestStatus
	CHANGE_JOB_FAILED               ListJobsRequestStatus
	CHILD_TRANSFER_STARTING         ListJobsRequestStatus
	CHILD_TRANSFER_STARTED          ListJobsRequestStatus
	CHILD_TRANSFER_COMPLETE         ListJobsRequestStatus
	CHILD_TRANSFER_FAILED           ListJobsRequestStatus
	RELEASE_CHILD_TRANSFER_STARTED  ListJobsRequestStatus
	RELEASE_CHILD_TRANSFER_COMPLETE ListJobsRequestStatus
}

func GetListJobsRequestStatusEnum() ListJobsRequestStatusEnum {
	return ListJobsRequestStatusEnum{
		CREATING: ListJobsRequestStatus{
			value: "CREATING",
		},
		CREATE_FAILED: ListJobsRequestStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: ListJobsRequestStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: ListJobsRequestStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: ListJobsRequestStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: ListJobsRequestStatus{
			value: "START_JOB_FAILED",
		},
		FULL_TRANSFER_STARTED: ListJobsRequestStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: ListJobsRequestStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: ListJobsRequestStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: ListJobsRequestStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: ListJobsRequestStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: ListJobsRequestStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: ListJobsRequestStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: ListJobsRequestStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		CHANGE_JOB_STARTED: ListJobsRequestStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: ListJobsRequestStatus{
			value: "CHANGE_JOB_FAILED",
		},
		CHILD_TRANSFER_STARTING: ListJobsRequestStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: ListJobsRequestStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: ListJobsRequestStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: ListJobsRequestStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: ListJobsRequestStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: ListJobsRequestStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
	}
}

func (c ListJobsRequestStatus) Value() string {
	return c.value
}

func (c ListJobsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestEngineType struct {
	value string
}

type ListJobsRequestEngineTypeEnum struct {
	ORACLE_TO_GAUSSDBV5        ListJobsRequestEngineType
	REDIS_TO_GAUSSREDIS        ListJobsRequestEngineType
	REDISCLUSTER_TO_GAUSSREDIS ListJobsRequestEngineType
}

func GetListJobsRequestEngineTypeEnum() ListJobsRequestEngineTypeEnum {
	return ListJobsRequestEngineTypeEnum{
		ORACLE_TO_GAUSSDBV5: ListJobsRequestEngineType{
			value: "oracle-to-gaussdbv5",
		},
		REDIS_TO_GAUSSREDIS: ListJobsRequestEngineType{
			value: "redis-to-gaussredis",
		},
		REDISCLUSTER_TO_GAUSSREDIS: ListJobsRequestEngineType{
			value: "rediscluster-to-gaussredis",
		},
	}
}

func (c ListJobsRequestEngineType) Value() string {
	return c.value
}

func (c ListJobsRequestEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestEngineType) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestNetType struct {
	value string
}

type ListJobsRequestNetTypeEnum struct {
	EIP ListJobsRequestNetType
	VPC ListJobsRequestNetType
	VPN ListJobsRequestNetType
}

func GetListJobsRequestNetTypeEnum() ListJobsRequestNetTypeEnum {
	return ListJobsRequestNetTypeEnum{
		EIP: ListJobsRequestNetType{
			value: "eip",
		},
		VPC: ListJobsRequestNetType{
			value: "vpc",
		},
		VPN: ListJobsRequestNetType{
			value: "vpn",
		},
	}
}

func (c ListJobsRequestNetType) Value() string {
	return c.value
}

func (c ListJobsRequestNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestNetType) UnmarshalJSON(b []byte) error {
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

type ListJobsRequestSortKey struct {
	value string
}

type ListJobsRequestSortKeyEnum struct {
	NAME          ListJobsRequestSortKey
	STATUS        ListJobsRequestSortKey
	CREATE_TIME   ListJobsRequestSortKey
	NET_TYPE      ListJobsRequestSortKey
	JOB_DIRECTION ListJobsRequestSortKey
	PAY_MODE      ListJobsRequestSortKey
}

func GetListJobsRequestSortKeyEnum() ListJobsRequestSortKeyEnum {
	return ListJobsRequestSortKeyEnum{
		NAME: ListJobsRequestSortKey{
			value: "name",
		},
		STATUS: ListJobsRequestSortKey{
			value: "status",
		},
		CREATE_TIME: ListJobsRequestSortKey{
			value: "create_time",
		},
		NET_TYPE: ListJobsRequestSortKey{
			value: "net_type",
		},
		JOB_DIRECTION: ListJobsRequestSortKey{
			value: "job_direction",
		},
		PAY_MODE: ListJobsRequestSortKey{
			value: "pay_mode",
		},
	}
}

func (c ListJobsRequestSortKey) Value() string {
	return c.value
}

func (c ListJobsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobsRequestSortKey) UnmarshalJSON(b []byte) error {
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
