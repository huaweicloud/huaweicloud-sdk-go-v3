package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListJobsRequest struct {

	// 请求语言类型。
	XLanguage *ListJobsRequestXLanguage `json:"X-Language,omitempty"`

	// 任务场景。取值： - migration：实时迁移。 - sync：实时同步。 - cloudDataGuard：实时灾备。
	JobType ListJobsRequestJobType `json:"job_type"`

	// 任务ID或名称。
	Name *string `json:"name,omitempty"`

	// 任务状态。取值： - CREATING：创建中。 - CREATE_FAILED：创建失败。 - CONFIGURATION：配置中。 - STARTJOBING：启动中。 - WAITING_FOR_START：等待启动中。 - START_JOB_FAILED：任务启动失败。 - FULL_TRANSFER_STARTED：全量迁移中，灾备场景为初始化。 - FULL_TRANSFER_FAILED：全量迁移失败，灾备场景为初始化失败。 - FULL_TRANSFER_COMPLETE：全量迁移完成，灾备场景为初始化完成。 - INCRE_TRANSFER_STARTED：增量迁移中，灾备场景为灾备中。 - INCRE_TRANSFER_FAILED：增量迁移失败，灾备场景为灾备异常。 - RELEASE_RESOURCE_STARTED：结束任务中。 - RELEASE_RESOURCE_FAILED：结束任务失败。 - RELEASE_RESOURCE_COMPLETE：已结束。 - CHANGE_JOB_STARTED：任务变更中。 - CHANGE_JOB_FAILED：任务变更失败。 - CHILD_TRANSFER_STARTING：子任务启动中。 - CHILD_TRANSFER_STARTED：子任务迁移中。 - CHILD_TRANSFER_COMPLETE：子任务迁移完成。 - CHILD_TRANSFER_FAILED：子任务迁移失败。 - RELEASE_CHILD_TRANSFER_STARTED：子任务结束中。 - RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束。 其中，异常状态可单独查询，也可以通过以下方式查询全部异常任务： CREATE_FAILED,START_JOB_FAILED,FULL_TRANSFER_FAILED,INCRE_TRANSFER_FAILED,RELEASE_RESOURCE_FAILED,CHANGE_JOB_FAILED,CHILD_TRANSFER_FAILED
	Status *ListJobsRequestStatus `json:"status,omitempty"`

	// 引擎类型。取值： - oracle-to-gaussdbv5：Oracle同步到GaussDB分布式版，实时同步场景使用。
	EngineType *ListJobsRequestEngineType `json:"engine_type,omitempty"`

	// 网络类型。取值： - eip：公网网络。 - vpc：VPC网络。 - vpn：VPN、专线网络。
	NetType *ListJobsRequestNetType `json:"net_type,omitempty"`

	// 企业项目ID。 缺省值：\"\"，表示查询所有企业项目任务。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 返回结果按该关键字排序，默认为“create_time”。
	SortKey *string `json:"sort_key,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为“desc”）。
	SortDir *string `json:"sort_dir,omitempty"`
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

type ListJobsRequestJobType struct {
	value string
}

type ListJobsRequestJobTypeEnum struct {
	MIGRATION        ListJobsRequestJobType
	SYNC             ListJobsRequestJobType
	CLOUD_DATA_GUARD ListJobsRequestJobType
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

type ListJobsRequestEngineType struct {
	value string
}

type ListJobsRequestEngineTypeEnum struct {
	ORACLE_TO_GAUSSDBV5 ListJobsRequestEngineType
}

func GetListJobsRequestEngineTypeEnum() ListJobsRequestEngineTypeEnum {
	return ListJobsRequestEngineTypeEnum{
		ORACLE_TO_GAUSSDBV5: ListJobsRequestEngineType{
			value: "oracle-to-gaussdbv5",
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
