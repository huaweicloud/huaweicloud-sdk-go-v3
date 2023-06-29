package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTasksRequest Request Object
type ListTasksRequest struct {

	// 迁移任务状态 READY:准备就绪 RUNNING:迁移中 SYNCING:同步中 MIGRATE_SUCCESS:迁移成功 MIGRATE_FAIL:迁移失败 ABORTING:中止中 ABORT:中止 DELETING::删除中 SYNC_F_ROLLBACKING:同步失败回滚中 SYNC_F_ROLLBACK_SUCCESS:同步失败回滚成功
	State *ListTasksRequestState `json:"state,omitempty"`

	// 任务的名称
	Name *string `json:"name,omitempty"`

	// 任务的ID
	Id *string `json:"id,omitempty"`

	// 源端服务器的ID
	SourceServerId *string `json:"source_server_id,omitempty"`

	// 每一页记录的任务数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 需要查询的企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}

type ListTasksRequestState struct {
	value string
}

type ListTasksRequestStateEnum struct {
	READY                   ListTasksRequestState
	RUNNING                 ListTasksRequestState
	SYNCING                 ListTasksRequestState
	MIGRATE_SUCCESS         ListTasksRequestState
	MIGRATE_FAIL            ListTasksRequestState
	ABORTING                ListTasksRequestState
	ABORT                   ListTasksRequestState
	DELETING                ListTasksRequestState
	SYNC_F_ROLLBACKING      ListTasksRequestState
	SYNC_F_ROLLBACK_SUCCESS ListTasksRequestState
}

func GetListTasksRequestStateEnum() ListTasksRequestStateEnum {
	return ListTasksRequestStateEnum{
		READY: ListTasksRequestState{
			value: "READY",
		},
		RUNNING: ListTasksRequestState{
			value: "RUNNING",
		},
		SYNCING: ListTasksRequestState{
			value: "SYNCING",
		},
		MIGRATE_SUCCESS: ListTasksRequestState{
			value: "MIGRATE_SUCCESS",
		},
		MIGRATE_FAIL: ListTasksRequestState{
			value: "MIGRATE_FAIL",
		},
		ABORTING: ListTasksRequestState{
			value: "ABORTING",
		},
		ABORT: ListTasksRequestState{
			value: "ABORT",
		},
		DELETING: ListTasksRequestState{
			value: "DELETING",
		},
		SYNC_F_ROLLBACKING: ListTasksRequestState{
			value: "SYNC_F_ROLLBACKING",
		},
		SYNC_F_ROLLBACK_SUCCESS: ListTasksRequestState{
			value: "SYNC_F_ROLLBACK_SUCCESS",
		},
	}
}

func (c ListTasksRequestState) Value() string {
	return c.value
}

func (c ListTasksRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestState) UnmarshalJSON(b []byte) error {
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
