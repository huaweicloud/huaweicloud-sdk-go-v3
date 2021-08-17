package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTasksRequest struct {
	// 迁移任务状态

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
	// 需要查询的企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := json.Marshal(o)
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
			value: "READY：准备就绪",
		},
		RUNNING: ListTasksRequestState{
			value: "RUNNING：迁移中",
		},
		SYNCING: ListTasksRequestState{
			value: "SYNCING：同步中",
		},
		MIGRATE_SUCCESS: ListTasksRequestState{
			value: "MIGRATE_SUCCESS：迁移成功",
		},
		MIGRATE_FAIL: ListTasksRequestState{
			value: "MIGRATE_FAIL：迁移失败",
		},
		ABORTING: ListTasksRequestState{
			value: "ABORTING：任务中止中",
		},
		ABORT: ListTasksRequestState{
			value: "ABORT：任务中止",
		},
		DELETING: ListTasksRequestState{
			value: "DELETING：删除中",
		},
		SYNC_F_ROLLBACKING: ListTasksRequestState{
			value: "SYNC_F_ROLLBACKING：同步失败回滚中",
		},
		SYNC_F_ROLLBACK_SUCCESS: ListTasksRequestState{
			value: "SYNC_F_ROLLBACK_SUCCESS：同步失败回滚成功",
		},
	}
}

func (c ListTasksRequestState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListTasksRequestState) UnmarshalJSON(b []byte) error {
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
