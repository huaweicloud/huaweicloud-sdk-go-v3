package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryJobStatusResp 批量查询任务状态返回体
type QueryJobStatusResp struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务状态。 - CREATING：创建中 - CREATE_FAILED：创建失败 - CONFIGURATION：配置中 - STARTJOBING：启动中 - WAITING_FOR_START：等待启动中 - START_JOB_FAILED：启动失败 - PAUSING：已暂停 - FULL_TRANSFER_STARTED：全量开始，灾备场景下为初始化 - FULL_TRANSFER_FAILED：全量失败，灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE：全量完成，灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED：增量开始，灾备场景下为灾备中 - INCRE_TRANSFER_FAILED：增量失败，灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED：结束任务中 - RELEASE_RESOURCE_FAILED：结束任务失败 - RELEASE_RESOURCE_COMPLETE：已结束 - REBUILD_NODE_STARTED：故障恢复中 - REBUILD_NODE_FAILED：故障恢复失败 - CHANGE_JOB_STARTED：任务变更中 - CHANGE_JOB_FAILED：任务变更失败 - DELETED：已删除 - CHILD_TRANSFER_STARTING：再编辑子任务启动中 - CHILD_TRANSFER_STARTED：再编辑子任务迁移中 - CHILD_TRANSFER_COMPLETE：再编辑子任务迁移完成 - CHILD_TRANSFER_FAILED：再编辑子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED：再编辑子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE：再编辑子任务已结束 - NODE_UPGRADE_START：升级开始 - NODE_UPGRADE_COMPLETE：升级完成 - NODE_UPGRADE_FAILED：升级失败
	Status *QueryJobStatusRespStatus `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o QueryJobStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryJobStatusResp struct{}"
	}

	return strings.Join([]string{"QueryJobStatusResp", string(data)}, " ")
}

type QueryJobStatusRespStatus struct {
	value string
}

type QueryJobStatusRespStatusEnum struct {
	CREATING                        QueryJobStatusRespStatus
	CREATE_FAILED                   QueryJobStatusRespStatus
	CONFIGURATION                   QueryJobStatusRespStatus
	STARTJOBING                     QueryJobStatusRespStatus
	WAITING_FOR_START               QueryJobStatusRespStatus
	START_JOB_FAILED                QueryJobStatusRespStatus
	PAUSING                         QueryJobStatusRespStatus
	FULL_TRANSFER_STARTED           QueryJobStatusRespStatus
	FULL_TRANSFER_FAILED            QueryJobStatusRespStatus
	FULL_TRANSFER_COMPLETE          QueryJobStatusRespStatus
	INCRE_TRANSFER_STARTED          QueryJobStatusRespStatus
	INCRE_TRANSFER_FAILED           QueryJobStatusRespStatus
	RELEASE_RESOURCE_STARTED        QueryJobStatusRespStatus
	RELEASE_RESOURCE_FAILED         QueryJobStatusRespStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobStatusRespStatus
	REBUILD_NODE_STARTED            QueryJobStatusRespStatus
	REBUILD_NODE_FAILED             QueryJobStatusRespStatus
	CHANGE_JOB_STARTED              QueryJobStatusRespStatus
	CHANGE_JOB_FAILED               QueryJobStatusRespStatus
	DELETED                         QueryJobStatusRespStatus
	CHILD_TRANSFER_STARTING         QueryJobStatusRespStatus
	CHILD_TRANSFER_STARTED          QueryJobStatusRespStatus
	CHILD_TRANSFER_COMPLETE         QueryJobStatusRespStatus
	CHILD_TRANSFER_FAILED           QueryJobStatusRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobStatusRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobStatusRespStatus
	NODE_UPGRADE_START              QueryJobStatusRespStatus
	NODE_UPGRADE_COMPLETE           QueryJobStatusRespStatus
	NODE_UPGRADE_FAILED             QueryJobStatusRespStatus
}

func GetQueryJobStatusRespStatusEnum() QueryJobStatusRespStatusEnum {
	return QueryJobStatusRespStatusEnum{
		CREATING: QueryJobStatusRespStatus{
			value: "CREATING",
		},
		CREATE_FAILED: QueryJobStatusRespStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: QueryJobStatusRespStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: QueryJobStatusRespStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: QueryJobStatusRespStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: QueryJobStatusRespStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: QueryJobStatusRespStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: QueryJobStatusRespStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: QueryJobStatusRespStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: QueryJobStatusRespStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: QueryJobStatusRespStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: QueryJobStatusRespStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobStatusRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: QueryJobStatusRespStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: QueryJobStatusRespStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: QueryJobStatusRespStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: QueryJobStatusRespStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: QueryJobStatusRespStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobStatusRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: QueryJobStatusRespStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: QueryJobStatusRespStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: QueryJobStatusRespStatus{
			value: "NODE_UPGRADE_FAILED",
		},
	}
}

func (c QueryJobStatusRespStatus) Value() string {
	return c.value
}

func (c QueryJobStatusRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobStatusRespStatus) UnmarshalJSON(b []byte) error {
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
