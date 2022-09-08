package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量查询任务状态返回体
type QueryJobStatusResp struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务状态
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
	FULL_TRANSFER_STARTED           QueryJobStatusRespStatus
	FULL_TRANSFER_FAILED            QueryJobStatusRespStatus
	FULL_TRANSFER_COMPLETE          QueryJobStatusRespStatus
	INCRE_TRANSFER_STARTED          QueryJobStatusRespStatus
	INCRE_TRANSFER_FAILED           QueryJobStatusRespStatus
	RELEASE_RESOURCE_STARTED        QueryJobStatusRespStatus
	RELEASE_RESOURCE_FAILED         QueryJobStatusRespStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobStatusRespStatus
	CHANGE_JOB_STARTED              QueryJobStatusRespStatus
	CHANGE_JOB_FAILED               QueryJobStatusRespStatus
	CHILD_TRANSFER_STARTING         QueryJobStatusRespStatus
	CHILD_TRANSFER_STARTED          QueryJobStatusRespStatus
	CHILD_TRANSFER_COMPLETE         QueryJobStatusRespStatus
	CHILD_TRANSFER_FAILED           QueryJobStatusRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobStatusRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobStatusRespStatus
}

func GetQueryJobStatusRespStatusEnum() QueryJobStatusRespStatusEnum {
	return QueryJobStatusRespStatusEnum{
		CREATING: QueryJobStatusRespStatus{
			value: "CREATING：创建中",
		},
		CREATE_FAILED: QueryJobStatusRespStatus{
			value: "CREATE_FAILED: 创建失败",
		},
		CONFIGURATION: QueryJobStatusRespStatus{
			value: "CONFIGURATION: 配置中",
		},
		STARTJOBING: QueryJobStatusRespStatus{
			value: "STARTJOBING: 启动中",
		},
		WAITING_FOR_START: QueryJobStatusRespStatus{
			value: "WAITING_FOR_START：等待启动中",
		},
		START_JOB_FAILED: QueryJobStatusRespStatus{
			value: "START_JOB_FAILED：任务启动失败",
		},
		FULL_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "FULL_TRANSFER_STARTED：全量迁移中，灾备场景为初始化",
		},
		FULL_TRANSFER_FAILED: QueryJobStatusRespStatus{
			value: "FULL_TRANSFER_FAILED：全量迁移失败，灾备场景为初始化失败",
		},
		FULL_TRANSFER_COMPLETE: QueryJobStatusRespStatus{
			value: "FULL_TRANSFER_COMPLETE：全量迁移完成，灾备场景为初始化完成",
		},
		INCRE_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "INCRE_TRANSFER_STARTED：增量迁移中，灾备场景为灾备中",
		},
		INCRE_TRANSFER_FAILED: QueryJobStatusRespStatus{
			value: "INCRE_TRANSFER_FAILED：增量迁移失败，灾备场景为灾备异常",
		},
		RELEASE_RESOURCE_STARTED: QueryJobStatusRespStatus{
			value: "RELEASE_RESOURCE_STARTED：结束任务中",
		},
		RELEASE_RESOURCE_FAILED: QueryJobStatusRespStatus{
			value: "RELEASE_RESOURCE_FAILED：结束任务失败",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobStatusRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE：已结束",
		},
		CHANGE_JOB_STARTED: QueryJobStatusRespStatus{
			value: "CHANGE_JOB_STARTED：任务变更中",
		},
		CHANGE_JOB_FAILED: QueryJobStatusRespStatus{
			value: "CHANGE_JOB_FAILED：任务变更失败",
		},
		CHILD_TRANSFER_STARTING: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_STARTING：子任务启动中",
		},
		CHILD_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_STARTED：子任务迁移中",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_COMPLETE：子任务迁移完成",
		},
		CHILD_TRANSFER_FAILED: QueryJobStatusRespStatus{
			value: "CHILD_TRANSFER_FAILED：子任务迁移失败",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobStatusRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED：子任务结束中",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobStatusRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE：子任务已结束",
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
